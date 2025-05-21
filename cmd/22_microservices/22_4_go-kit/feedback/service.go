package feedback

import "context"

// Service interface
// define methods that you want to expose and transport layer can use to interact with your service implementatio
type Service interface {
	AddFeedback(ctx context.Context, feedback Feedback) (Feedback, error)
	GetFeedback(ctx context.Context, id string) (Feedback, error)
	GetAllFeedback(ctx context.Context) ([]Feedback, error)
}
