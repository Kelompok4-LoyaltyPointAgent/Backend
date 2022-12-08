package constant

type MidtransStatusEnum string

const (
	MidtransStatusSettlement MidtransStatusEnum = "settlement"
	MidtransStatusPending    MidtransStatusEnum = "pending"
	MidtransStatusExpire     MidtransStatusEnum = "expire"
	MidtransStatusCancel     MidtransStatusEnum = "cancel"
	MidtransStatusDeny       MidtransStatusEnum = "deny"
	MidtransStatusRefund     MidtransStatusEnum = "refund"
)
