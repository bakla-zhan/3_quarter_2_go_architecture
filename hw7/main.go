package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"shop/middlewares"
	"shop/pkg/mailsender"
	"shop/pkg/tgbot"
	"shop/repository"
	"shop/service"
)

func main() {
	var token = flag.String("token", "", "telegram bot token")
	var chatID = flag.Int64("chat", 0, "telegram chat id")
	var server = flag.String("server", "smtp.mail.ru", "notification smtp server address")
	var sender = flag.String("sender", "", "notification sender email address")
	var password = flag.String("password", "", "notification sender password")

	flag.Parse()

	tg, err := tgbot.NewTelegramAPI(*token, *chatID)
	if err != nil {
		log.Fatal("Unable to init telegram bot")
	}

	ms, err := mailsender.NewMailAPI(*server, *sender, *password)
	if err != nil {
		log.Fatal("Unable to init mail sender")
	}

	db := repository.NewMapDB()

	service := service.NewService(tg, ms, db)
	handler := &shopHandler{
		service: service,
		db:      db,
	}

	router := mux.NewRouter()
	router.Use(middlewares.RequestIDMiddleware)

	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")

	router.HandleFunc("/order", handler.createOrderHandler).Methods("POST")
	router.HandleFunc("/order/{id}", handler.getOrderHandler).Methods("GET")

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	log.Println("the server is running...")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
