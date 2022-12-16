package favorites_service

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/favorites_repository"
)

type FavoritesService interface {
	FindAll(claims *helper.JWTCustomClaims) (*[]response.FavoritesResponse, error)
	Create(payload payload.FavoritesPayload, id string) (*response.FavoritesResponse, error)
	Delete(userID string, productID string) error
}

type favoritesService struct {
	favoritesRepository favorites_repository.FavoritesRepository
}

func NewFavoritesService(favoritesRepository favorites_repository.FavoritesRepository) FavoritesService {
	return &favoritesService{favoritesRepository}
}

func (s *favoritesService) FindAll(claims *helper.JWTCustomClaims) (*[]response.FavoritesResponse, error) {
	var args []any
	var favorites []models.Favorites
	var err error
	if claims.Role == constant.UserRoleAdmin.String() {
		favorites, err = s.favoritesRepository.FindAll()
		if err != nil {
			return nil, err
		}
	} else {
		args = append(args, claims.ID.String())
		favorites, err = s.favoritesRepository.FindAll(args...)
		if err != nil {
			return nil, err
		}
	}

	return response.NewFavoritesListResponse(favorites), nil
}

func (s *favoritesService) Create(payload payload.FavoritesPayload, id string) (*response.FavoritesResponse, error) {
	favorites := models.Favorites{
		UserID:    uuid.MustParse(id),
		ProductID: uuid.MustParse(payload.ProductID),
	}
	favorites, err := s.favoritesRepository.Create(favorites)
	if err != nil {
		return nil, err
	}
	return response.NewFavoritesResponse(favorites), nil
}

func (s *favoritesService) Delete(userID string, productID string) error {
	err := s.favoritesRepository.Delete(userID, productID)
	if err != nil {
		return err
	}
	return nil
}
