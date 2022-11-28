package response

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/models"
)

type ProductResponse struct {
	ID               uuid.UUID  `json:"id"`
	Name             string     `json:"name"`
	Type             string     `json:"type"`
	Provider         string     `json:"provider"`
	Price            uint       `json:"price"`
	PricePoints      uint       `json:"price_points"`
	RewardPoints     uint       `json:"reward_points"`
	Stock            uint       `json:"stock"`
	Recommended      bool       `json:"recommended"`
	ProductPictureID *uuid.UUID `json:"product_picture_id,omitempty"`
}

//Product With Credit Response
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

func NewProductsWithCreditsResponse(products []models.Product, credits []models.Credit) *[]ProductWithCreditResponse {
	var response []ProductWithCreditResponse
	for i := range products {
		credits[i].Product = nil
		response = append(response, *NewProductWithCreditResponse(products[i], credits[i]))
	}
	return &response
}

//Product With Package Response
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

func NewProductsWithPackagesResponse(products []models.Product, packages []models.Packages) *[]ProductWithPackagesResponse {
	var response []ProductWithPackagesResponse
	for i := range products {
		packages[i].Product = nil
		response = append(response, *NewProductWithPackagesResponse(products[i], packages[i]))
	}
	return &response
}

func NewProductResponse(product models.Product) *ProductResponse {
	return &ProductResponse{
		ID:               product.ID,
		Name:             product.Name,
		Type:             product.Type,
		Provider:         product.Provider,
		Price:            product.Price,
		PricePoints:      product.PricePoints,
		RewardPoints:     product.RewardPoints,
		Stock:            product.Stock,
		Recommended:      product.Recommended,
		ProductPictureID: product.ProductPictureID,
	}
}

func NewProductsResponse(products []models.Product) *[]ProductResponse {
	var response []ProductResponse
	for _, product := range products {
		response = append(response, *NewProductResponse(product))
	}
	return &response
}
