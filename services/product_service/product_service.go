package product_service

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_repository"
)

type ProductService interface {
	FindAll() (*[]response.ProductResponse, error)
	CreateProduct(payload payload.ProductPayload) (*response.ProductResponse, error)
}

type productService struct {
	repository product_repository.ProductRepository
}

func NewProductService(repository product_repository.ProductRepository) ProductService {
	return &productService{repository}
}

func (s *productService) FindAll() (*[]response.ProductResponse, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return response.NewProductsResponse(products), nil
}

func (s *productService) CreateProduct(payload payload.ProductPayload) (*response.ProductResponse, error) {
	product := models.Product{
		Name:         payload.Name,
		Type:         payload.Type,
		Provider:     payload.Provider,
		Price:        payload.Price,
		PricePoints:  payload.PricePoints,
		RewardPoints: payload.RewardPoints,
		Stock:        payload.Stock,
		Recommended:  payload.Recommended,
	}

	if payload.ProductPictureID != "" {
		id, err := uuid.Parse(payload.ProductPictureID)
		if err != nil {
			return nil, err
		}
		product.ProductPictureID = &id
	}

	product, err := s.repository.Create(product)
	if err != nil {
		return nil, err
	}

	return response.NewProductResponse(product), nil
}
