package faq_service

import (
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/faq_repository"
)

type FAQService interface {
	FindAll(query any, args ...any) (*[]response.FAQResponse, error)
	FindByID(id any) (*response.FAQResponse, error)
	Create(payload payload.FAQPayload) (*response.FAQResponse, error)
	Update(payload payload.FAQPayload, id any) (*response.FAQResponse, error)
	Delete(id any) error
}

type faqService struct {
	repository faq_repository.FAQRepository
}

func NewFAQService(repository faq_repository.FAQRepository) FAQService {
	return &faqService{repository}
}

func (s *faqService) FindAll(query any, args ...any) (*[]response.FAQResponse, error) {
	faqs, err := s.repository.FindAll(query, args...)
	if err != nil {
		return nil, err
	}

	return response.NewFAQsResponse(faqs), nil
}

func (s *faqService) FindByID(id any) (*response.FAQResponse, error) {
	faq, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return response.NewFAQResponse(faq), nil
}

func (s *faqService) Create(payload payload.FAQPayload) (*response.FAQResponse, error) {
	faq, err := s.repository.Create(models.FAQ{
		Category: payload.Category,
		Question: payload.Question,
		Answer:   payload.Answer,
	})
	if err != nil {
		return nil, err
	}

	return response.NewFAQResponse(faq), nil
}

func (s *faqService) Update(payload payload.FAQPayload, id any) (*response.FAQResponse, error) {
	faq, err := s.repository.Update(models.FAQ{
		Category: payload.Category,
		Question: payload.Question,
		Answer:   payload.Answer,
	}, id)
	if err != nil {
		return nil, err
	}

	return response.NewFAQResponse(faq), nil
}

func (s *faqService) Delete(id any) error {
	return s.repository.Delete(id)
}
