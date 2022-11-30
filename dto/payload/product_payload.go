package payload

import "mime/multipart"

type ProductPayload struct {
	Name           string                `form:"name" validate:"required"`
	Provider       string                `form:"provider" validate:"required"`
	Price          uint                  `form:"price" validate:"required"`
	PricePoints    uint                  `form:"price_points" validate:"required"`
	RewardPoints   uint                  `form:"reward_points" validate:"required"`
	Stock          uint                  `form:"stock" validate:"required"`
	Recommended    bool                  `form:"recommended"`
	ProductPicture *multipart.FileHeader `form:"product_picture"`
	Description    string                `form:"description" validate:"required"`
	TermOfService  string                `form:"term_of_service" validate:"required"`
}

type ProductWithCreditPayload struct {
	ProductPayload
	ActivePeriod int `form:"active_period" validate:"required"`
	Amount       int `form:"amount" validate:"required"`
}

type ProductWithPackagesPayload struct {
	ProductPayload
	ActivePeriod int `form:"active_period" validate:"required"`
	Internet     int `form:"internet" validate:"required"`
	Call         int `form:"call" validate:"required"`
}
