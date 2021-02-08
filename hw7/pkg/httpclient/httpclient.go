package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"shop/middlewares"
	"shop/models"
	"time"
)

type WebAPI1C interface {
	SendOrderTo1CWebAPI(ctx context.Context, order *models.Order) error
}

type webAPI1C struct {
	uri string
}

func NewWebAPI1C(uri string) (*webAPI1C, error) {
	return &webAPI1C{
		uri: uri,
	}, nil
}

func (w *webAPI1C) SendOrderTo1CWebAPI(ctx context.Context, order *models.Order) error {
	var id string
	cli := &http.Client{
		Timeout: 10 * time.Second,
	}
	b, err := json.Marshal(order)
	if err != nil {
		log.Println(err)
		return err
	}
	body := bytes.NewBuffer(b)
	req, err := http.NewRequest(
		"POST",
		w.uri,
		body,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	id, ok := ctx.Value(middlewares.CtxRequestIDKey).(string)
	if !ok {
		return errors.New("unable to get request ID from Context")
	}
	req.Header.Add(middlewares.CtxRequestIDKey, id)

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	if (resp.StatusCode / 100) != 2 {
		return errors.New("unable to sent order to 1C")
	}
	defer resp.Body.Close()
	return nil
}
