package payload

type FAQPayload struct {
	Category string `json:"category" validate:"required"`
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
}
