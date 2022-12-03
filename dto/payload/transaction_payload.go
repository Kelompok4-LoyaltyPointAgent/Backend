package payload

import "github.com/kelompok4-loyaltypointagent/backend/constant"

type TransactionPayload struct {
	UserID    string                         `json:"user_id"`
	ProductID string                         `json:"product_id" validate:"required"`
	Amount    float64                        `json:"amount"`
	Method    string                         `json:"method"`
	Number    string                         `json:"number" validate:"required"`
	Email     string                         `json:"email" validate:"required,email"`
	Status    constant.TransactionStatusEnum `json:"status"`
	Type      constant.TransactionTypeEnum   `json:"type" validate:"required,eq=Redeem|eq=Purchase|eq=Cashout"`
}
