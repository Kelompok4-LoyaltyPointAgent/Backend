package product_service

import (
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/credit_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/packages_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_picture_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_repository"
)

type ProductService interface {
	//Credit
	FindAllWithCredits() (*[]response.ProductWithCreditResponse, error)
	FindByIDWithCredit(id any) (*response.ProductWithCreditResponse, error)
	CreateProductWithCredit(payload payload.ProductWithCreditPayload) (*response.ProductWithCreditResponse, error)
	UpdateProductWithCredit(payload payload.ProductWithCreditPayload, id any) (*response.ProductWithCreditResponse, error)
	DeleteProductWithCredit(id any) error

	//Packages
	FindAllWithPackages() (*[]response.ProductWithPackagesResponse, error)
	FindByIDWithPackages(id any) (*response.ProductWithPackagesResponse, error)
	CreateProductWithPackages(payload payload.ProductWithPackagesPayload) (*response.ProductWithPackagesResponse, error)
	UpdateProductWithPackages(payload payload.ProductWithPackagesPayload, id any) (*response.ProductWithPackagesResponse, error)
	DeleteProductWithPackages(id any) error
}

type productService struct {
	productRepository        product_repository.ProductRepository
	creditRepository         credit_repository.CreditRepository
	packagesRepository       packages_repository.PackagesRepository
	productPictureRepository product_picture_repository.ProductPictureRepository
}

func NewProductService(productRepository product_repository.ProductRepository, creditRepository credit_repository.CreditRepository, packagesRepository packages_repository.PackagesRepository, productPictureRepository product_picture_repository.ProductPictureRepository) *productService {
	return &productService{productRepository, creditRepository, packagesRepository, productPictureRepository}
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
		Name:           payload.Name,
		Type:           "Credit",
		Provider:       payload.Provider,
		Price:          payload.Price,
		PricePoints:    payload.PricePoints,
		RewardPoints:   payload.RewardPoints,
		Stock:          payload.Stock,
		Recommended:    payload.Recommended,
		Description:    payload.Description,
		TermsOfService: payload.TermOfService,
	}

	if payload.ProductPicture != nil {
		fileName, buf, err := helper.OpenFileFromMultipartForm(payload.ProductPicture)
		if err != nil {
			return nil, err
		}
		productPicture, err := s.productPictureRepository.FindByName(fileName)
		if err != nil {
			url := helper.UploadFileToFirebase(*buf, fileName)
			createProductPicture, err := s.productPictureRepository.Create(models.ProductPicture{
				Name: fileName,
				Url:  url,
			})
			if err != nil {
				return nil, err
			}
			product.ProductPictureID = &createProductPicture.ID
		} else {
			product.ProductPictureID = &productPicture.ID
		}
	}

	product, err := s.productRepository.Create(product)
	if err != nil {
		return nil, err
	}

	credit, err := s.creditRepository.Create(models.Credit{
		ProductID:    &product.ID,
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
		Name:           payload.Name,
		Type:           "Credit",
		Provider:       payload.Provider,
		Price:          payload.Price,
		PricePoints:    payload.PricePoints,
		RewardPoints:   payload.RewardPoints,
		Stock:          payload.Stock,
		Recommended:    payload.Recommended,
		Description:    payload.Description,
		TermsOfService: payload.TermOfService,
	}

	if payload.ProductPicture != nil {
		fileName, buf, err := helper.OpenFileFromMultipartForm(payload.ProductPicture)
		if err != nil {
			return nil, err
		}
		productPicture, err := s.productPictureRepository.FindByName(fileName)
		if err != nil {
			url := helper.UploadFileToFirebase(*buf, fileName)
			createProductPicture, err := s.productPictureRepository.Create(models.ProductPicture{
				Name: fileName,
				Url:  url,
			})
			if err != nil {
				return nil, err
			}
			product.ProductPictureID = &createProductPicture.ID
		} else {
			product.ProductPictureID = &productPicture.ID
		}
	}

	product, err := s.productRepository.Update(product, id)
	if err != nil {
		return nil, err
	}

	credit, err := s.creditRepository.UpdateByProductID(models.Credit{
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

func (s *productService) FindAllWithPackages() (*[]response.ProductWithPackagesResponse, error) {
	products, err := s.productRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var packages []models.Packages
	for _, product := range products {
		pack, err := s.packagesRepository.FindByProductID(product.ID)
		if err != nil {
			continue
		}
		packages = append(packages, pack)
	}

	return response.NewProductsWithPackagesResponse(products, packages), nil
}

func (s *productService) FindByIDWithPackages(id any) (*response.ProductWithPackagesResponse, error) {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	pack, err := s.packagesRepository.FindByProductID(id)
	if err != nil {
		return nil, err
	}

	return response.NewProductWithPackagesResponse(product, pack), nil
}

func (s *productService) CreateProductWithPackages(payload payload.ProductWithPackagesPayload) (*response.ProductWithPackagesResponse, error) {
	product := models.Product{
		Name:           payload.Name,
		Type:           "Packages",
		Provider:       payload.Provider,
		Price:          payload.Price,
		PricePoints:    payload.PricePoints,
		RewardPoints:   payload.RewardPoints,
		Stock:          payload.Stock,
		Recommended:    payload.Recommended,
		Description:    payload.Description,
		TermsOfService: payload.TermOfService,
	}

	if payload.ProductPicture != nil {
		fileName, buf, err := helper.OpenFileFromMultipartForm(payload.ProductPicture)
		if err != nil {
			return nil, err
		}
		productPicture, err := s.productPictureRepository.FindByName(fileName)
		if err != nil {
			url := helper.UploadFileToFirebase(*buf, fileName)
			createProductPicture, err := s.productPictureRepository.Create(models.ProductPicture{
				Name: fileName,
				Url:  url,
			})
			if err != nil {
				return nil, err
			}
			product.ProductPictureID = &createProductPicture.ID
		} else {
			product.ProductPictureID = &productPicture.ID
		}
	}

	product, err := s.productRepository.Create(product)
	if err != nil {
		return nil, err
	}

	pack, err := s.packagesRepository.Create(models.Packages{
		ProductID: &product.ID,
		Internet:  payload.Internet,
		Call:      payload.Call,
	})
	if err != nil {
		return nil, err
	}

	return response.NewProductWithPackagesResponse(product, pack), nil
}

func (s *productService) UpdateProductWithPackages(payload payload.ProductWithPackagesPayload, id any) (*response.ProductWithPackagesResponse, error) {
	product := models.Product{
		Name:           payload.Name,
		Type:           "Packages",
		Provider:       payload.Provider,
		Price:          payload.Price,
		PricePoints:    payload.PricePoints,
		RewardPoints:   payload.RewardPoints,
		Stock:          payload.Stock,
		Recommended:    payload.Recommended,
		Description:    payload.Description,
		TermsOfService: payload.TermOfService,
	}

	if payload.ProductPicture != nil {
		fileName, buf, err := helper.OpenFileFromMultipartForm(payload.ProductPicture)
		if err != nil {
			return nil, err
		}
		productPicture, err := s.productPictureRepository.FindByName(fileName)
		if err != nil {
			url := helper.UploadFileToFirebase(*buf, fileName)
			createProductPicture, err := s.productPictureRepository.Create(models.ProductPicture{
				Name: fileName,
				Url:  url,
			})
			if err != nil {
				return nil, err
			}
			product.ProductPictureID = &createProductPicture.ID
		} else {
			product.ProductPictureID = &productPicture.ID
		}
	}

	product, err := s.productRepository.Update(product, id)
	if err != nil {
		return nil, err
	}

	pack, err := s.packagesRepository.UpdateByProductID(models.Packages{
		Internet: payload.Internet,
		Call:     payload.Call,
	}, id)
	if err != nil {
		return nil, err
	}

	return response.NewProductWithPackagesResponse(product, pack), nil
}

func (s *productService) DeleteProductWithPackages(id any) error {
	if err := s.packagesRepository.DeleteByProductID(id); err != nil {
		return err
	}

	if err := s.productRepository.DeleteByID(id); err != nil {
		return err
	}

	return nil
}
