package constant

type TransactionTypeEnum string

const (
	TransactionTypePurchase TransactionTypeEnum = "purchase"
	TransactionTypeRedeem   TransactionTypeEnum = "redeem"
)

type TransactionPaymentMethodEnum string

const (
	TransactionPaymentMethodGopay TransactionPaymentMethodEnum = "gopay"
	TransactionPaymentMethodQris  TransactionPaymentMethodEnum = "qris"
	TransactionPaymentMethodBank  TransactionPaymentMethodEnum = "bank_transfer"
)

type XenditStatusEnum string

const (
	XenditStatusPending   XenditStatusEnum = "PENDING"
	XenditStatusVoided    XenditStatusEnum = "VOIDED"
	XenditStatusCompleted XenditStatusEnum = "COMPLETED"
	XenditStatusFailed    XenditStatusEnum = "FAILED"
	XenditStatusExpired   XenditStatusEnum = "EXPIRED"
)

type MidtransStatusEnum string

const (
	MidtransStatusSettlement MidtransStatusEnum = "settlement"
	MidtransStatusPending    MidtransStatusEnum = "pending"
	MidtransStatusExpire     MidtransStatusEnum = "expire"
	MidtransStatusCancel     MidtransStatusEnum = "cancel"
	MidtransStatusDeny       MidtransStatusEnum = "deny"
	MidtransStatusRefund     MidtransStatusEnum = "refund"
)
