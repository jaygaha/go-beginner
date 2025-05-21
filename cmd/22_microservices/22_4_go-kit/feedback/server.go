package feedback

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer creates a new HTTP server.
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(forceJSONmiddleware)

	// register endpoints
	r.Methods("POST").Path("/feedbacks").Handler(httptransport.NewServer(
		endpoints.AddFeedbackEndpoint,
		decodeAddFeedbackRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/feedbacks/{id}").Handler(httptransport.NewServer(
		endpoints.GetFeedbackEndpoint,
		decodeGetFeedbackRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/feedbacks").Handler(httptransport.NewServer(
		endpoints.GetAllFeedbackEndpoint,
		decodeGetAllFeedbackRequest,
		encodeResponse,
	))

	return r
}

// forceJSONmiddleware is force json header
func forceJSONmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
