package chatbot

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ReciveWebhookEndpoint endpoint.Endpoint
}

// MakeReciveWebhookEndpoint
func MakeReciveWebhookEndpoint(srv LinebotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(webhookRequest)
		d, err := srv.ReciveWebhook(ctx, req.Message)
		if err != nil {
			return getResponse{d, err.Error()}, nil
		}
		return getResponse{d, ""}, nil
	}
}
