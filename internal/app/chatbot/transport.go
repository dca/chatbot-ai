package chatbot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type webhookResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

type webhookRequest struct {
	Message string `json:"message"`
}

func decodeWebhookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req webhookRequest
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

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeWebhookResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
