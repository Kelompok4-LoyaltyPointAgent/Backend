package response

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"github.com/kelompok4-loyaltypointagent/backend/models"
)

type ProductResponse struct {
	ID             uuid.UUID                `json:"id"`
	Name           string                   `json:"name"`
	Description    string                   `json:"description"`
	Type           constant.ProductTypeEnum `json:"type"`
	Provider       string                   `json:"provider"`
	Price          uint                     `json:"price"`
	PricePoints    uint                     `json:"price_points"`
	RewardPoints   uint                     `json:"reward_points"`
	Stock          uint                     `json:"stock"`
	Recommended    bool                     `json:"recommended"`
	ProductPicture ProductPicture           `json:"product_picture,omitempty"`
}

type ProductPicture struct {
	ID   uuid.UUID                       `json:"id"`
	Name string                          `json:"name"`
	Url  string                          `json:"url"`
	Type constant.ProductPictureTypeEnum `json:"type"`
}

// Product With Credit Response
type ProductWithCreditResponse struct {
	ProductResponse
	Credit CreditResponse `json:"credit,omitempty"`
}

func NewProductWithCreditResponse(product models.Product, credit models.Credit) *ProductWithCreditResponse {
	credit.ProductID = nil
	return &ProductWithCreditResponse{
		ProductResponse: *NewProductResponse(product),
		Credit:          *NewCreditResponse(credit),
	}
}

func NewProductsWithCreditsResponse(credits []models.Credit) *[]ProductWithCreditResponse {
	var response []ProductWithCreditResponse
	for _, cred := range credits {
		if cred.Product.Name != "" {
			response = append(response, *NewProductWithCreditResponse(cred.Product, cred))
		}
	}
	return &response
}

// Product With Package Response
type ProductWithPackagesResponse struct {
	ProductResponse
	Package PackagesResponse `json:"package,omitempty"`
}

func NewProductWithPackagesResponse(product models.Product, packages models.Packages) *ProductWithPackagesResponse {
	packages.ProductID = nil
	return &ProductWithPackagesResponse{
		ProductResponse: *NewProductResponse(product),
		Package:         *NewPackagesResponse(packages),
	}
}

func NewProductsWithPackagesResponse(packages []models.Packages) *[]ProductWithPackagesResponse {
	var response []ProductWithPackagesResponse
	for _, pack := range packages {
		if pack.Product.Name != "" {
			response = append(response, *NewProductWithPackagesResponse(pack.Product, pack))
		}
	}
	return &response
}

func NewProductResponse(product models.Product) *ProductResponse {
	var productPicture ProductPicture
	if product.ProductPicture != nil {
		productPicture = ProductPicture{
			ID:   *product.ProductPictureID,
			Name: product.ProductPicture.Name,
			Url:  product.ProductPicture.Url,
		}
	}

	productResponse := ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Type:         product.Type,
		Provider:     product.Provider,
		Price:        product.Price,
		PricePoints:  product.PricePoints,
		RewardPoints: product.RewardPoints,
		Description:  product.Description,
		Stock:        product.Stock,
		Recommended:  product.Recommended,
	}

	if product.ProductPicture != nil {
		productResponse.ProductPicture = productPicture
	}

	return &productResponse
}

func NewProductsResponse(products []models.Product) *[]ProductResponse {
	var response []ProductResponse
	for _, product := range products {
		response = append(response, *NewProductResponse(product))
	}
	return &response
}
