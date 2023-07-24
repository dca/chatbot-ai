package message

import "context"

// MessageService provides operations on strings.
type MessageService interface {
	HandlerTextMessage(ctx context.Context, message string) (string, error)
}

// messageService is a concrete implementation of MessageService
type messageService struct{}

func (messageService) HandlerTextMessage(ctx context.Context, message string) (string, error) {
	return "ok", nil
}

func NewMessageService() MessageService {
	return messageService{}
}
