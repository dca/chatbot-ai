package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/dca/chatbot-ai/internal/app/message"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8081", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := message.NewMessageService()
	errChan := make(chan error)

	endpoints := message.Endpoints{
		HandleMessageEndpoint: message.MakeHandleMessageEndpoint(srv),
	}

	go func() {
		log.Println("chatbot is listening on port:", *httpAddr)
		handler := message.NewHttpServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
