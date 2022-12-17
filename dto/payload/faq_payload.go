package payload

type FAQPayload struct {
	Category string `json:"category" validate:"required,min=1,max=50"`
	Question string `json:"question" validate:"required,min=1,max=300"`
	Answer   string `json:"answer" validate:"required,min=1,max=1000"`
}
