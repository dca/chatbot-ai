package message

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

// MessageService provides operations on strings.
type MessageService interface {
	HandlerTextMessage(ctx context.Context, message string) (string, error)
}

// messageService is a concrete implementation of MessageService
type messageService struct{}

func (messageService) HandlerTextMessage(ctx context.Context, message string) (string, error) {
	resp, err := sendMessageToOpenai(message)

	if err != nil {
		fmt.Printf("sendMessageToOpenai error: %v\n", err)
		return "", err
	}

	return resp, nil
}

func NewMessageService() MessageService {
	return messageService{}
}

func sendMessageToOpenai(message string) (string, error) {
	apikey := viper.GetString("OPENAI_API_KEY")
	fmt.Printf("apikey: %s\n", apikey)

	client := openai.NewClient(apikey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}
