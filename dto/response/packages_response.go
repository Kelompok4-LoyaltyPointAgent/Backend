package response

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/models"
)

type PackagesResponse struct {
	ID            uuid.UUID        `json:"id"`
	ProductID     *uuid.UUID       `json:"product_id,omitempty"`
	Product       *ProductResponse `json:"product,omitempty"`
	ActivePeriod  int              `json:"active_period"`
	TotalInternet float64          `json:"total_internet"`
	MainInternet  float64          `json:"main_internet"`
	NightInternet float64          `json:"night_internet"`
	SocialMedia   float64          `json:"social_media"`
	Call          int              `json:"call"`
	SMS           int              `json:"sms"`
	TermOfService string           `json:"term_of_service"`
}

func NewPackagesResponse(packages models.Packages) *PackagesResponse {
	response := &PackagesResponse{
		ID:            packages.ID,
		ProductID:     packages.ProductID,
		ActivePeriod:  packages.ActivePeriod,
		TotalInternet: packages.TotalInternet,
		MainInternet:  packages.MainInternet,
		NightInternet: packages.NightInternet,
		SocialMedia:   packages.SocialMedia,
		Call:          packages.Call,
		SMS:           packages.SMS,
		TermOfService: packages.TermsOfService,
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
