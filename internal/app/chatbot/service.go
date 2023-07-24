package chatbot

import "context"

// LinebotService provides operations on strings.
type LinebotService interface {
	ReciveWebhook(ctx context.Context, message string) (string, error)
}

// linebotService is a concrete implementation of LinebotService
type linebotService struct{}

func (linebotService) ReciveWebhook(ctx context.Context, message string) (string, error) {
	return "ok", nil
}

func NewLinebotService() LinebotService {
	return linebotService{}
}
