package message

import (
	"testing"

	"github.com/dca/chatbot-ai/internal/pkg/config"
)

func TestSendMessageToOpenai(t *testing.T) {
	_, err := config.LoadConfig("./../../../")
	if err != nil {
		t.Errorf("LoadConfig error: %v", err)
	}

	sendMessageToOpenai("hello")
}
