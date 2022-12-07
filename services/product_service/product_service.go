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
	FindByProviderWithCredit(provider string) (*[]response.ProductWithCreditResponse, error)
	FindByRecommendedWithCredit() (*[]response.ProductWithCreditResponse, error)
	CreateProductWithCredit(payload payload.ProductWithCreditPayload) (*response.ProductWithCreditResponse, error)
	UpdateProductWithCredit(payload payload.ProductWithCreditPayload, id any) (*response.ProductWithCreditResponse, error)
	DeleteProductWithCredit(id any) error

	//Packages
	FindAllWithPackages() (*[]response.ProductWithPackagesResponse, error)
	FindByIDWithPackages(id any) (*response.ProductWithPackagesResponse, error)
	FindByProviderWithPackages(provider string) (*[]response.ProductWithPackagesResponse, error)
	FindByRecommendedWithPackages() (*[]response.ProductWithPackagesResponse, error)
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
	credits, err := s.creditRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return response.NewProductsWithCreditsResponse(credits), nil
}

func (s *productService) FindByIDWithCredit(id any) (*response.ProductWithCreditResponse, error) {

	credit, err := s.creditRepository.FindByProductID(id)
	if err != nil {
		return nil, err
	}

	return response.NewProductWithCreditResponse(credit.Product, credit), nil
}

func (s *productService) FindByProviderWithCredit(provider string) (*[]response.ProductWithCreditResponse, error) {
	credits, err := s.creditRepository.FindByProvider(provider)
	if err != nil {
		return nil, err
	}

	return response.NewProductsWithCreditsResponse(credits), nil
}

func (s *productService) FindByRecommendedWithCredit() (*[]response.ProductWithCreditResponse, error) {
	credits, err := s.creditRepository.FindByRecommended()
	if err != nil {
		return nil, err
	}

	return response.NewProductsWithCreditsResponse(credits), nil
}

func (s *productService) CreateProductWithCredit(payload payload.ProductWithCreditPayload) (*response.ProductWithCreditResponse, error) {
	product := models.Product{
		Name:         payload.Name,
		Description:  payload.Description,
		Type:         "Credit",
		Provider:     payload.Provider,
		Price:        payload.Price,
		PricePoints:  payload.PricePoints,
		RewardPoints: payload.RewardPoints,
		Stock:        payload.Stock,
	}

	if payload.Recommended != nil {
		product.Recommended = *payload.Recommended
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
		Call:         payload.Call,
		SMS:          payload.SMS,
	})
	if err != nil {
		return nil, err
	}

	return response.NewProductWithCreditResponse(product, credit), nil
}

func (s *productService) UpdateProductWithCredit(payload payload.ProductWithCreditPayload, id any) (*response.ProductWithCreditResponse, error) {
	product := models.Product{
		Name:         payload.Name,
		Description:  payload.Description,
		Type:         "Credit",
		Provider:     payload.Provider,
		Price:        payload.Price,
		PricePoints:  payload.PricePoints,
		RewardPoints: payload.RewardPoints,
		Stock:        payload.Stock,
	}

	if payload.Recommended != nil {
		//Update Recommended
		err := s.productRepository.SetBooleanRecommended(id, *payload.Recommended)
		if err != nil {
			return nil, err
		}
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
			product.ProductPicture = &createProductPicture
		} else {
			product.ProductPictureID = &productPicture.ID
			product.ProductPicture = &productPicture
		}
	}

	product, err := s.productRepository.Update(product, id)
	if err != nil {
		return nil, err
	}

	credit, err := s.creditRepository.UpdateByProductID(models.Credit{
		ActivePeriod: payload.ActivePeriod,
		Amount:       payload.Amount,
		Call:         payload.Call,
		SMS:          payload.SMS,
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
	pack, err := s.packagesRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return response.NewProductsWithPackagesResponse(pack), nil
}

func (s *productService) FindByIDWithPackages(id any) (*response.ProductWithPackagesResponse, error) {

	pack, err := s.packagesRepository.FindByProductID(id)
	if err != nil {
		return nil, err
	}

	return response.NewProductWithPackagesResponse(pack.Product, pack), nil
}

func (s *productService) FindByProviderWithPackages(provider string) (*[]response.ProductWithPackagesResponse, error) {
	packages, err := s.packagesRepository.FindByProvider(provider)
	if err != nil {
		return nil, err
	}

	return response.NewProductsWithPackagesResponse(packages), nil
}

func (s *productService) FindByRecommendedWithPackages() (*[]response.ProductWithPackagesResponse, error) {
	packages, err := s.packagesRepository.FindByRecommended()
	if err != nil {
		return nil, err
	}

	return response.NewProductsWithPackagesResponse(packages), nil
}

func (s *productService) CreateProductWithPackages(payload payload.ProductWithPackagesPayload) (*response.ProductWithPackagesResponse, error) {
	product := models.Product{
		Name:         payload.Name,
		Description:  payload.Description,
		Type:         "Packages",
		Provider:     payload.Provider,
		Price:        payload.Price,
		PricePoints:  payload.PricePoints,
		RewardPoints: payload.RewardPoints,
		Stock:        payload.Stock,
	}

	if payload.Recommended != nil {
		product.Recommended = *payload.Recommended
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
			product.ProductPicture = &createProductPicture
		} else {
			product.ProductPictureID = &productPicture.ID
			product.ProductPicture = &productPicture
		}
	}

	product, err := s.productRepository.Create(product)
	if err != nil {
		return nil, err
	}

	pack, err := s.packagesRepository.Create(models.Packages{
		ProductID:     &product.ID,
		ActivePeriod:  payload.ActivePeriod,
		TotalInternet: payload.TotalInternet,
		MainInternet:  payload.MainInternet,
		NightInternet: payload.NightInternet,
		SocialMedia:   payload.SocialMedia,
		Call:          payload.Call,
		SMS:           payload.SMS,
		Description:   payload.Description,
	})
	if err != nil {
		return nil, err
	}

	return response.NewProductWithPackagesResponse(product, pack), nil
}

func (s *productService) UpdateProductWithPackages(payload payload.ProductWithPackagesPayload, id any) (*response.ProductWithPackagesResponse, error) {
	product := models.Product{
		Name:         payload.Name,
		Description:  payload.Description,
		Type:         "Packages",
		Provider:     payload.Provider,
		Price:        payload.Price,
		PricePoints:  payload.PricePoints,
		RewardPoints: payload.RewardPoints,
		Stock:        payload.Stock,
	}

	if payload.Recommended != nil {
		//Update Recommended
		err := s.productRepository.SetBooleanRecommended(id, *payload.Recommended)
		if err != nil {
			return nil, err
		}
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
		ActivePeriod:  payload.ActivePeriod,
		TotalInternet: payload.TotalInternet,
		MainInternet:  payload.MainInternet,
		NightInternet: payload.NightInternet,
		SocialMedia:   payload.SocialMedia,
		Call:          payload.Call,
		SMS:           payload.SMS,
		Description:   payload.Description,
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
