package response

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/models"
)

type FAQResponse struct {
	ID       uuid.UUID `json:"id"`
	Category string    `json:"category"`
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
}

func NewFAQResponse(faq models.FAQ) *FAQResponse {
	response := &FAQResponse{
		ID:       faq.ID,
		Category: faq.Category,
		Question: faq.Question,
		Answer:   faq.Answer,
	}
	return response
}

func NewFAQsResponse(faqs []models.FAQ) *[]FAQResponse {
	var response []FAQResponse
	for _, faq := range faqs {
		response = append(response, *NewFAQResponse(faq))
	}
	return &response
}
