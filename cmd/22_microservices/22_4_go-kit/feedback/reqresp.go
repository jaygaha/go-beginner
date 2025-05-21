package feedback

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	AddFeedbackRequest struct {
		Message string `json:"message"`
	}
	AddFeedbackResponse struct {
		Feedback Feedback `json:"feedback"`
		Err      error    `json:"err,omitempty"`
	}
	GetFeedbackRequest struct {
		ID string `json:"id"`
	}
	GetFeedbackResponse struct {
		Feedback Feedback `json:"feedback"`
		Err      error    `json:"err,omitempty"`
	}
	GetAllFeedbackRequest  struct{}
	GetAllFeedbackResponse struct {
		Feedback []Feedback `json:"feedbacks"`
		Err      error      `json:"err,omitempty"`
	}
)

// encodeResponse encodes the response to the client.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response any) error {
	return json.NewEncoder(w).Encode(response)
}

// decodeAddFeedbackRequest decodes the request from the client.
func decodeAddFeedbackRequest(ctx context.Context, r *http.Request) (any, error) {
	var req AddFeedbackRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	// validate request
	if req.Message == "" {
		return nil, fmt.Errorf("message is required")
	}

	return req, nil
}

// decodeGetFeedbackRequest decodes the request from the client.
func decodeGetFeedbackRequest(ctx context.Context, r *http.Request) (any, error) {
	var req GetFeedbackRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	req.ID = vars["id"]

	// validate request
	if req.ID == "" {
		return nil, fmt.Errorf("id is required")
	}

	return req, nil
}

// decodeGetAllFeedbackRequest decodes the request from the client.
func decodeGetAllFeedbackRequest(ctx context.Context, r *http.Request) (any, error) {
	return GetAllFeedbackRequest{}, nil
}
