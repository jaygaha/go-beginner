package feedback

import (
	"context"
	"time"
)

type Feedback struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type FeedbackRepository interface {
	AddFeedback(ctx context.Context, feedback Feedback) (Feedback, error)
	GetFeedback(ctx context.Context, id string) (Feedback, error)
	GetAllFeedback(ctx context.Context) ([]Feedback, error)
}
