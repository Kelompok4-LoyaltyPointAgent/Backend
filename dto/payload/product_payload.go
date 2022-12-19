package payload

import "mime/multipart"

type ProductPayload struct {
	Name           string                `form:"name" validate:"required,min=1,max=100"`
	Description    string                `form:"description" validate:"required,min=1,max=500"`
	Provider       string                `form:"provider" validate:"required,min=1,max=50"`
	Price          uint                  `form:"price"`
	PricePoints    uint                  `form:"price_points"`
	RewardPoints   uint                  `form:"reward_points"`
	Stock          uint                  `form:"stock"`
	Recommended    *bool                 `form:"recommended"`
	ProductPicture *multipart.FileHeader `form:"product_picture"`
}

type ProductWithCreditPayload struct {
	ProductPayload
	ActivePeriod int `form:"active_period"`
	Amount       int `form:"amount"`
	Call         int `form:"call"`
	SMS          int `form:"sms"`
}

type ProductWithPackagesPayload struct {
	ProductPayload
	ActivePeriod       int     `form:"active_period"`
	TotalInternet      float64 `form:"total_internet"`
	MainInternet       float64 `form:"main_internet"`
	NightInternet      float64 `form:"night_internet"`
	SocialMedia        float64 `form:"social_media"`
	Call               int     `form:"call"`
	SMS                int     `form:"sms"`
	PackageDescription string  `form:"package_description" validate:"required,min=1,max=1000"`
}
