package message

import (
	"context"
	"encoding/json"
	"net/http"
)

type handlerTextMessageRequest struct {
	Message string `json:"message"`
}

type handlerTextMessageResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

func decodeHandleMessageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req handlerTextMessageRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeHandleMessageResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
