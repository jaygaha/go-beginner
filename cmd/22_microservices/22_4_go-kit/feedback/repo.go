package feedback

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
)

var RepoErr = "unable to handle request"

// in-memory feedback repository
type repo struct {
	feedbacks map[string]Feedback
	logger    log.Logger
}

// NewRepo creates a new feedback repository.
func NewRepo(logger log.Logger) FeedbackRepository {
	return &repo{
		feedbacks: make(map[string]Feedback),
		logger:    log.With(logger, "repo", "feedback"),
	}
}

// AddFeedback adds a new feedback to the repository.
func (r *repo) AddFeedback(ctx context.Context, feedback Feedback) (Feedback, error) {
	// add feedback to repository
	// add logic to add into the database, json file, etc.
	r.feedbacks[feedback.ID] = feedback

	return feedback, nil
}

// GetFeedback retrieves a feedback from the repository.
func (r *repo) GetFeedback(ctx context.Context, id string) (Feedback, error) {
	// retrieve feedback from repository
	// add logic to retrieve from the database, json file, etc.
	feedback, ok := r.feedbacks[id]
	if !ok {
		return Feedback{}, fmt.Errorf("feedback not found")
	}

	return feedback, nil
}

// GetAllFeedback retrieves all feedback from the repository.
func (r *repo) GetAllFeedback(ctx context.Context) ([]Feedback, error) {
	// retrieve all feedback from repository
	// add logic to retrieve from the database, json file, etc.
	var feedbacks []Feedback
	for _, feedback := range r.feedbacks {
		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}
