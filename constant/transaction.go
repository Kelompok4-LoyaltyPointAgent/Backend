package constant

type TransactionTypeEnum string

const (
	TransactionTypePurchase TransactionTypeEnum = "Purchase"
	TransactionTypeRedeem   TransactionTypeEnum = "Redeem"
)

type TransactionStatusEnum string

const (
	TransactionStatusSuccess TransactionStatusEnum = "Success"
	TransactionStatusPending TransactionStatusEnum = "Pending"
	TransactionStatusFailed  TransactionStatusEnum = "Failed"
)

type TransactionPaymentMethodEnum string

const (
	TransactionPaymentMethodGopay TransactionPaymentMethodEnum = "Gopay"
	TransactionPaymentMethodQris  TransactionPaymentMethodEnum = "QRIS"
	TransactionPaymentMethodBank  TransactionPaymentMethodEnum = "Bank"
)
