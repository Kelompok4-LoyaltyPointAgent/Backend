package payload

type PackagesPayload struct {
	ProductID    string `json:"product_id" validate:"required"`
	ActivePeriod int    `json:"active_period" validate:"required"`
	Internet     int    `json:"internet" validate:"required"`
	Call         int    `json:"call" validate:"required"`
}
