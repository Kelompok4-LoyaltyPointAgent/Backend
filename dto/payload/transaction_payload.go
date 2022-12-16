package payload

import "github.com/kelompok4-loyaltypointagent/backend/constant"

type TransactionPayload struct {
	UserID    string                         `json:"user_id"`
	ProductID string                         `json:"product_id"`
	Amount    float64                        `json:"amount"`
	Method    string                         `json:"method"`
	Number    string                         `json:"number" validate:"required,min=1,max=20"`
	Email     string                         `json:"email" validate:"email"`
	Status    constant.TransactionStatusEnum `json:"status"`
	Type      constant.TransactionTypeEnum   `json:"type" validate:"required,eq=Redeem|eq=Purchase|eq=Cashout"`
}

type TransactionUpdatePayload struct {
	Amount float64                        `json:"amount" validate:"required"`
	Status constant.TransactionStatusEnum `json:"status" validate:"required,eq=Pending|eq=Success|eq=Failed"`
	Type   constant.TransactionTypeEnum   `json:"type" validate:"required,eq=Redeem|eq=Purchase|eq=Cashout"`
}
