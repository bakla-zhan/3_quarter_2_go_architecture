package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	totalReq       int64
	successReq     int64
	avarageReqTime int64
)

// Job описание задания для воркера
type Job struct {
	method string
	URI    string
	body   io.Reader
}

// Worker описание структуры воркера
type Worker struct {
	wg          *sync.WaitGroup
	jobChan     <-chan *Job
	successChan chan<- struct{}
}

// NewWorker функция для создания нового экземпляра воркера
func NewWorker(wg *sync.WaitGroup, jobChan <-chan *Job, successChan chan<- struct{}) *Worker {
	return &Worker{
		wg:          wg,
		jobChan:     jobChan,
		successChan: successChan,
	}
}

// Handle обработчик задания
func (w *Worker) Handle() {
	defer w.wg.Done()
	var requests, totalTime int64
	success := struct{}{}
	for job := range w.jobChan {
		code, duration, err := sendHTTPRequest(job.method, job.URI, job.body)
		if err != nil {
			log.Println(err)
		}
		if (code/100 == 2) || (code/100 == 3) {
			w.successChan <- success
		}
		requests++
		totalTime += duration.Microseconds()
		avarageReqTime = totalTime / requests
	}
}

// sendHTTPRequest функция для отправки http-запросов
func sendHTTPRequest(method, URI string, body io.Reader) (int, time.Duration, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		method,
		URI,
		body,
	)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	start := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(start)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, duration, nil
}

func successCount(successChan <-chan struct{}) {
	for range successChan {
		successReq++
	}
}

func printStat(duration time.Duration) {
	t := time.After(duration)
	for {
		select {
		case <-t:
			return
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Print("\033[H\033[2J")
			fmt.Printf("текущее среднее время отклика запроса - %.2f миллисекунд\n", float32(avarageReqTime)/1000)
		}
	}
}

func main() {
	addr := flag.String("address", "http://localhost:8080", "server address for payload test")
	method := flag.String("method", "GET", "http request method")
	inputBody := flag.String("body", "", "http request body")
	threads := flag.Int("threads", 2, "treads number of ddos")
	inputDuration := flag.String("duration", "5", "test duration in seconds") // конечно можно было использвать flag.Duration(), но тогда длительность теста пришлось бы указывать в наносекундах, что недружелюбно для пользователя))

	flag.Parse()

	// преобразование некоторых флагов в нужный нам формат -->
	body := strings.NewReader(*inputBody)
	duration, err := time.ParseDuration(fmt.Sprint(*inputDuration, "s"))
	if err != nil {
		log.Fatal("incorrect duration input format", err)
	}
	durationFloat, err := strconv.ParseFloat(*inputDuration, 64)
	if err != nil {
		log.Fatal("incorrect duration input format", err)
	}
	// <--

	wg := &sync.WaitGroup{}
	jobChan := make(chan *Job)
	successChan := make(chan struct{})
	for i := 0; i < *threads; i++ {
		worker := NewWorker(wg, jobChan, successChan)
		wg.Add(1)
		go worker.Handle()
	}

	go successCount(successChan)

	func() {
		t := time.After(duration)
		go printStat(duration)
		for {
			select {
			case <-t:
				return
			default:
				jobChan <- &Job{
					method: *method,
					URI:    *addr,
					body:   body,
				}
				totalReq++
			}
		}
	}()

	close(jobChan)
	wg.Wait()
	fmt.Print("\033[H\033[2J")
	fmt.Printf("RPS: %.2f запр/сек\nСреднее время отклика запроса: %.2f миллисекунд\nОбщее кол-во запросов: %v\nУспешных запросов: %v\nУспешных запросов %%: %.2f%%\n", float64(totalReq)/durationFloat, float32(avarageReqTime)/1000, totalReq, successReq, float64(successReq)/float64(totalReq)*100)
}
