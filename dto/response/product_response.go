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
