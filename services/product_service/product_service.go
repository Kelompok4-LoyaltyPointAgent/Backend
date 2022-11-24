package product_service

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/credit_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_repository"
)

type ProductService interface {
	FindAllWithCredits() (*[]response.ProductWithCreditResponse, error)
	FindByIDWithCredit(id any) (*response.ProductWithCreditResponse, error)
	CreateProductWithCredit(payload payload.ProductWithCreditPayload) (*response.ProductWithCreditResponse, error)
	UpdateProductWithCredit(payload payload.ProductWithCreditPayload, id any) (*response.ProductWithCreditResponse, error)
	DeleteProductWithCredit(id any) error
}

type productService struct {
	productRepository product_repository.ProductRepository
	creditRepository  credit_repository.CreditRepository
}

func NewProductService(productRepository product_repository.ProductRepository, creditRepository credit_repository.CreditRepository) ProductService {
	return &productService{productRepository, creditRepository}
}

func (s *productService) FindAllWithCredits() (*[]response.ProductWithCreditResponse, error) {
	products, err := s.productRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var credits []models.Credit
	for _, product := range products {
		credit, err := s.creditRepository.FindByProductID(product.ID)
		if err != nil {
			continue
		}
		credits = append(credits, credit)
	}

	return response.NewProductsWithCreditsResponse(products, credits), nil
}

func (s *productService) FindByIDWithCredit(id any) (*response.ProductWithCreditResponse, error) {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	credit, err := s.creditRepository.FindByProductID(id)
	if err != nil {
		return nil, err
	}

	return response.NewProductWithCreditResponse(product, credit), nil
}

func (s *productService) CreateProductWithCredit(payload payload.ProductWithCreditPayload) (*response.ProductWithCreditResponse, error) {
	product := models.Product{
		Name:         payload.Name,
		Type:         "Credit",
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

	product, err := s.productRepository.Create(product)
	if err != nil {
		return nil, err
	}

	credit, err := s.creditRepository.Create(models.Credit{
		ProductID:    &product.ID,
		Description:  payload.Description,
		ActivePeriod: payload.ActivePeriod,
		Amount:       payload.Amount,
	})
	if err != nil {
		return nil, err
	}

	return response.NewProductWithCreditResponse(product, credit), nil
}

func (s *productService) UpdateProductWithCredit(payload payload.ProductWithCreditPayload, id any) (*response.ProductWithCreditResponse, error) {
	product := models.Product{
		Name:         payload.Name,
		Type:         "Credit",
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

	product, err := s.productRepository.Update(product, id)
	if err != nil {
		return nil, err
	}

	credit, err := s.creditRepository.UpdateByProductID(models.Credit{
		Description:  payload.Description,
		ActivePeriod: payload.ActivePeriod,
		Amount:       payload.Amount,
	}, id)
	if err != nil {
		return nil, err
	}

	return response.NewProductWithCreditResponse(product, credit), nil
}

func (s *productService) DeleteProductWithCredit(id any) error {
	if err := s.creditRepository.DeleteByProductID(id); err != nil {
		return err
	}

	if err := s.productRepository.DeleteByID(id); err != nil {
		return err
	}

	return nil
}
