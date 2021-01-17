package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"shop/repository"
	"shop/service"
)

func main() {
	isDebug := flag.Bool("mode", true, "runs app in debug mode")
	flag.Parse()

	os.Setenv("WEB_SERVER_ADDR", "127.0.0.1:8080")

	webAddr, ok := os.LookupEnv("WEB_SERVER_ADDR")
	if !ok {
		log.Fatal("WEB_SERVER_ADDR env not set")
	}

	handler := &shopHandler{}
	if *isDebug {
		rep := repository.NewMapDB()
		handler.db = rep
		handler.service = service.NewService(rep)
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")

	srv := &http.Server{
		Addr: webAddr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
