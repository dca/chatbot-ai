package message

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {

	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/message").Handler(httptransport.NewServer(
		endpoints.HandleMessageEndpoint,
		decodeHandleMessageRequest,
		encodeHandleMessageResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
