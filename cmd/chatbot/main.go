package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/dca/chatbot-ai/internal/app/chatbot"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := chatbot.NewLinebotService()
	errChan := make(chan error)

	endpoints := chatbot.Endpoints{
		ReciveWebhookEndpoint: chatbot.MakeReciveWebhookEndpoint(srv),
	}

	go func() {
		log.Println("chatbot is listening on port:", *httpAddr)
		handler := chatbot.NewHttpServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
