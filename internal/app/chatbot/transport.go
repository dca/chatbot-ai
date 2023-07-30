package chatbot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

type webhookResponse struct {
	Messages []string `json:"messages"`
	Errs     []string `json:"errs,omitempty"`
}

type webhookRequest struct {
	Destination string `json:"destination"`
	Events      []struct {
		Type    string `json:"type"`
		Message struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"message"`
		WebhookEventId  string `json:"webhookEventId"`
		DeliveryContext struct {
			IsRedelivery bool `json:"isRedelivery"`
		} `json:"deliveryContext"`
		Timestamp int64 `json:"timestamp"`
		Source    struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		ReplyToken string `json:"replyToken"`
		Mode       string `json:"mode"`
	} `json:"events"`
}

func decodeWebhookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var err error

	// only for debug in local
	if os.Getenv("ENV") == "local" {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		fmt.Println(string(bodyBytes))
	}

	lintSec := viper.GetString("LINE_CHANNEL_SECRET")
	lintToken := viper.GetString("LINE_CHANNEL_TOKEN")
	log.Println("token: ", lintSec, lintToken)

	bot, err := linebot.New(viper.GetString("LINE_CHANNEL_SECRET"), viper.GetString("LINE_CHANNEL_TOKEN"))
	events, err := bot.ParseRequest(r)

	if err != nil {
		log.Println("ParseRequest Error: ", err.Error())
		return nil, err
	}
	return events, nil
}

func encodeWebhookResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
