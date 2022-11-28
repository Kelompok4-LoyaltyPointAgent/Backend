package packages_service

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/packages_repository"
)

type PackagesService interface {
	FindAll() (*[]response.PackagesResponse, error)
	CreatePackages(payload payload.PackagesPayload) (*response.PackagesResponse, error)
}

type packagesService struct {
	repository packages_repository.PackagesRepository
}

func NewPackagesService(repository packages_repository.PackagesRepository) *packagesService {
	return &packagesService{repository}
}

func (s *packagesService) FindAll() (*[]response.PackagesResponse, error) {
	packages, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return response.NewPackagesResponses(packages), nil
}

func (s *packagesService) CreatePackages(payload payload.PackagesPayload) (*response.PackagesResponse, error) {
	packages := models.Packages{
		ActivePeriod: payload.ActivePeriod,
		Internet:     payload.Internet,
		Call:         payload.Call,
	}

	if payload.ProductID != "" {
		id, err := uuid.Parse(payload.ProductID)
		if err != nil {
			return nil, err
		}
		packages.ProductID = &id
	}

	packages, err := s.repository.Create(packages)
	if err != nil {
		return nil, err
	}

	return response.NewPackagesResponse(packages), nil
}


