package feedback

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints stuct defines the endpoint for the feedback service.
type Endpoints struct {
	AddFeedbackEndpoint    endpoint.Endpoint
	GetFeedbackEndpoint    endpoint.Endpoint
	GetAllFeedbackEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct that wraps the provided service.
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		AddFeedbackEndpoint:    makeAddFeedbackEndpoint(s),
		GetFeedbackEndpoint:    makeGetFeedbackEndpoint(s),
		GetAllFeedbackEndpoint: makeGetAllFeedbackEndpoint(s),
	}
}

// makeAddFeedbackEndpoint returns an endpoint for the AddFeedback method.
func makeAddFeedbackEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(AddFeedbackRequest)
		feedbak := Feedback{
			Message: req.Message,
		}
		feedback, err := s.AddFeedback(ctx, feedbak)

		return AddFeedbackResponse{Feedback: feedback, Err: err}, nil
	}
}

// makeGetFeedbackEndpoint returns an endpoint for the GetFeedback method.
func makeGetFeedbackEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(GetFeedbackRequest)
		feedback, err := s.GetFeedback(ctx, req.ID)

		return GetFeedbackResponse{Feedback: feedback, Err: err}, nil
	}
}

// makeGetAllFeedbackEndpoint returns an endpoint for the GetAllFeedback method.
func makeGetAllFeedbackEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		feedback, err := s.GetAllFeedback(ctx)

		return GetAllFeedbackResponse{Feedback: feedback, Err: err}, nil
	}
}
