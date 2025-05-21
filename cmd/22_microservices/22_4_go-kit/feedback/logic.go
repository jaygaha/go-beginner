package feedback

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type service struct {
	repo   FeedbackRepository
	logger log.Logger
}

// NewService creates a new feedback service.
func NewService(repo FeedbackRepository, logger log.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

// AddFeedback adds a new feedback.
func (s *service) AddFeedback(ctx context.Context, feedback Feedback) (Feedback, error) {
	logger := log.With(s.logger, "method", "AddFeedback")
	logger.Log("msg", "adding feedback")

	feedback.ID = "F" + time.Now().Format("20060102150405")
	feedback.CreatedAt = time.Now()

	feedback, err := s.repo.AddFeedback(ctx, feedback)
	if err != nil {
		level.Error(logger).Log("err", err)
		return Feedback{}, err
	}

	return feedback, nil
}

// GetFeedback retrieves a feedback by ID.
func (s *service) GetFeedback(ctx context.Context, id string) (Feedback, error) {
	logger := log.With(s.logger, "method", "GetFeedback")
	logger.Log("msg", "getting feedback", "id", id)

	feedback, err := s.repo.GetFeedback(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return Feedback{}, err
	}

	logger.Log("msg", "feedback retrieved", "feedback", feedback)

	return feedback, nil
}

// GetAllFeedback retrieves all feedback.
func (s *service) GetAllFeedback(ctx context.Context) ([]Feedback, error) {
	logger := log.With(s.logger, "method", "GetAllFeedback")
	logger.Log("msg", "getting all feedback")

	feedbacks, err := s.repo.GetAllFeedback(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return []Feedback{}, err
	}

	logger.Log("msg", "feedbacks retrieved", "feedbacks", feedbacks)

	return feedbacks, nil
}
