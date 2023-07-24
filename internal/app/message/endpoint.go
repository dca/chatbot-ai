package message

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	HandleMessageEndpoint endpoint.Endpoint
}

// MakeHandleMessageEndpoint
func MakeHandleMessageEndpoint(srv MessageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(handlerTextMessageRequest)
		d, err := srv.HandlerTextMessage(ctx, req.Message)
		if err != nil {
			return handlerTextMessageResponse{d, err.Error()}, nil
		}
		return handlerTextMessageResponse{d, ""}, nil
	}
}
