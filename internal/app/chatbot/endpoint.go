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
		d, errs := srv.ReciveWebhook(ctx, req)
		errStrings := make([]string, len(errs))
		for i, err := range errs {
			errStrings[i] = err.Error()
		}
		return webhookResponse{d, errStrings}, nil
	}
}
