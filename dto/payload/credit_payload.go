package payload

type CreditPayload struct {
	ProductID    string `json:"product_id" validate:"required"`
	ActivePeriod int    `json:"active_period" validate:"required"`
	Amount       int    `json:"amount" validate:"required"`
}
