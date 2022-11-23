package payload

type ProductPayload struct {
	Name             string `json:"name" validate:"required"`
	Type             string `json:"type" validate:"required"`
	Provider         string `json:"provider" validate:"required"`
	Price            uint   `json:"price" validate:"required"`
	PricePoints      uint   `json:"price_points" validate:"required"`
	RewardPoints     uint   `json:"reward_points" validate:"required"`
	Stock            uint   `json:"stock" validate:"required"`
	Recommended      bool   `json:"recommended"`
	ProductPictureID string `json:"product_picture_id"`
}

type ProductCreditPayload struct {
	ProductPayload
	CreditPayload
}
