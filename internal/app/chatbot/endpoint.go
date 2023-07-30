package chatbot

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Endpoints struct {
	ReciveWebhookEndpoint endpoint.Endpoint
}

// MakeReciveWebhookEndpoint
func MakeReciveWebhookEndpoint(srv LinebotService) endpoint.Endpoint {
	return func(ctx context.Context, events interface{}) (interface{}, error) {

		d, errs := srv.ReciveWebhook(ctx, events.([]*linebot.Event))
		errStrings := make([]string, len(errs))
		for i, err := range errs {
			errStrings[i] = err.Error()
		}
		return webhookResponse{d, errStrings}, nil
	}
}
