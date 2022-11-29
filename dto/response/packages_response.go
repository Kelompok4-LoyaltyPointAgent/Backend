package response

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/models"
)

type PackagesResponse struct {
	ID           uuid.UUID        `json:"id"`
	ProductID    *uuid.UUID       `json:"product_id,omitempty"`
	Product      *ProductResponse `json:"product,omitempty"`
	ActivePeriod int              `json:"active_period"`
	Internet     int              `json:"internet"`
	Call         int              `json:"call"`
}

func NewPackagesResponse(packages models.Packages) *PackagesResponse {
	response := &PackagesResponse{
		ID:           packages.ID,
		ProductID:    packages.ProductID,
		ActivePeriod: packages.ActivePeriod,
		Internet:     packages.Internet,
		Call:         packages.Call,
	}
	return response
}

func NewPackagesResponses(packages []models.Packages) *[]PackagesResponse {
	var response []PackagesResponse
	for _, packages := range packages {
		response = append(response, *NewPackagesResponse(packages))
	}
	return &response
}
