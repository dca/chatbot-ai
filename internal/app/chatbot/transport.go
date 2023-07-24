package chatbot

import (
	"context"
	"encoding/json"
	"net/http"
)

type getRequest struct{}

type getResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

type webhookRequest struct {
	Message string `json:"message"`
}

type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}

func decodeWebhookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req webhookRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeWebhookResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
