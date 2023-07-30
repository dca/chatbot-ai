package chatbot

import (
	"context"
	"log"

	"github.com/dca/chatbot-ai/internal/app/message"
)

// LinebotService provides operations on strings.
type LinebotService interface {
	ReciveWebhook(ctx context.Context, req webhookRequest) ([]string, []error)
}

// linebotService is a concrete implementation of LinebotService
type linebotService struct{}

func (linebotService) ReciveWebhook(ctx context.Context, req webhookRequest) ([]string, []error) {

	var responseSlice []string
	var errorSlice []error
	messageService := message.NewMessageService()

	for _, event := range req.Events {
		if event.Type == "message" && event.Message.Type == "text" {
			// replyToken := event.ReplyToken
			messageText := event.Message.Text
			resp, err := messageService.HandlerTextMessage(ctx, messageText)

			log.Println(resp)
			if err != nil {
				errorSlice = append(errorSlice, err)
			}
			responseSlice = append(responseSlice, resp)
		}
	}
	return responseSlice, errorSlice
}

func NewLinebotService() LinebotService {
	return linebotService{}
}
