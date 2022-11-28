package credit_service

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/credit_repository"
)

type CreditService interface {
	FindAll() (*[]response.CreditResponse, error)
	CreateCredit(payload payload.CreditPayload) (*response.CreditResponse, error)
}

type creditService struct {
	repository credit_repository.CreditRepository
}

func NewCreditService(repository credit_repository.CreditRepository) CreditService {
	return &creditService{repository}
}

func (s *creditService) FindAll() (*[]response.CreditResponse, error) {
	credits, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return response.NewCreditsResponse(credits), nil
}

func (s *creditService) CreateCredit(payload payload.CreditPayload) (*response.CreditResponse, error) {
	credit := models.Credit{
		ActivePeriod: payload.ActivePeriod,
		Amount:       payload.Amount,
	}

	if payload.ProductID != "" {
		id, err := uuid.Parse(payload.ProductID)
		if err != nil {
			return nil, err
		}
		credit.ProductID = &id
	}

	credit, err := s.repository.Create(credit)
	if err != nil {
		return nil, err
	}

	return response.NewCreditResponse(credit), nil
}
