package response

import "github.com/google/uuid"

type FAQResponse struct {
	ID       uuid.UUID `json:"id"`
	Category string    `json:"category"`
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
}
