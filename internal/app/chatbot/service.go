package chatbot

import (
	"context"
	"log"

	"github.com/dca/chatbot-ai/internal/app/message"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

// LinebotService provides operations on strings.
type LinebotService interface {
	ReciveWebhook(ctx context.Context, req []*linebot.Event) ([]string, []error)
}

// linebotService is a concrete implementation of LinebotService
type linebotService struct{}

func (linebotService) ReciveWebhook(ctx context.Context, events []*linebot.Event) ([]string, []error) {
	bot, err := linebot.New(viper.GetString("LINE_CHANNEL_SECRET"), viper.GetString("LINE_CHANNEL_TOKEN"))
	if err != nil {
		return nil, []error{err}
	}

	var responseSlice []string
	var errorSlice []error
	messageService := message.NewMessageService()

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				resp, err := messageService.HandlerTextMessage(ctx, message.Text)
				if err != nil {
					log.Println("messageService.HandlerTextMessage Error: ", err.Error())

					errorSlice = append(errorSlice, err)
					continue
				}

				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(resp)).Do(); err != nil {
					log.Println("ReplyMessage Error: ", err.Error())

					errorSlice = append(errorSlice, err)
					continue
				}

				responseSlice = append(responseSlice, resp)

			}
		}
	}
	return responseSlice, errorSlice
}

func NewLinebotService() LinebotService {
	return linebotService{}
}
