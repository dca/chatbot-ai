package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dca/chatbot-ai/internal/app/chatbot"
	"github.com/dca/chatbot-ai/internal/pkg/config"
	"github.com/spf13/viper"
)

func main() {
	_, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	srv := chatbot.NewLinebotService()
	errChan := make(chan error)

	endpoints := chatbot.Endpoints{
		ReciveWebhookEndpoint: chatbot.MakeReciveWebhookEndpoint(srv),
	}

	go func() {
		log.Printf("chatbot is listening on port: %s", viper.GetString("CHATBOT_PORT"))
		handler := chatbot.NewHttpServer(ctx, endpoints)
		errChan <- http.ListenAndServe(viper.GetString("CHATBOT_PORT"), handler)
	}()

	log.Fatalln(<-errChan)
}
