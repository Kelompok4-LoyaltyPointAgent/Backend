package feedback_service

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/feedback_repository"
)

type FeedbackService interface {
	FindAll() (*[]response.FeedbackResponse, error)
	FindByID(id any) (*response.FeedbackResponse, error)
	Create(feedback payload.FeedbackPayload, id string) (*response.FeedbackResponse, error)
}

type feedbackService struct {
	repository feedback_repository.FeedbackRepository
}

func NewFeedbackService(repository feedback_repository.FeedbackRepository) FeedbackService {
	return &feedbackService{repository}
}

func (s *feedbackService) FindAll() (*[]response.FeedbackResponse, error) {
	feedback, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return response.NewFeedbackResponseList(feedback), nil
}

func (s *feedbackService) FindByID(id any) (*response.FeedbackResponse, error) {
	feedback, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return response.NewFeedbackResponse(feedback), nil
}

func (s *feedbackService) Create(payload payload.FeedbackPayload, id string) (*response.FeedbackResponse, error) {
	feedback := models.Feedbacks{
		UserID:               uuid.MustParse(id),
		IsInformationHelpful: payload.IsInformationHelpful,
		IsArticleHelpful:     payload.IsArticleHelpful,
		IsArticleEasyToFind:  payload.IsArticleEasyToFind,
		IsDesignGood:         payload.IsDesignGood,
		Review:               payload.Review,
	}

	feedback, err := s.repository.Create(feedback)
	if err != nil {
		return nil, err
	}

	return response.NewFeedbackResponse(feedback), nil
}
