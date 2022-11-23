package payload

type CreditPayload struct {
	ProductID    string `json:"product_id" validate:"required"`
	Description  string `json:"description" validate:"required"`
	ActivePeriod int    `json:"active_period" validate:"required"`
	Amount       int    `json:"amount" validate:"required"`
}
