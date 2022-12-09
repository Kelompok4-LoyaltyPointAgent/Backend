package constant

type TransactionTypeEnum string

const (
	TransactionTypePurchase TransactionTypeEnum = "Purchase"
	TransactionTypeRedeem   TransactionTypeEnum = "Redeem"
	TransactionTypeCashout  TransactionTypeEnum = "Cashout"
)

type TransactionStatusEnum string

const (
	TransactionStatusSuccess TransactionStatusEnum = "Success"
	TransactionStatusPending TransactionStatusEnum = "Pending"
	TransactionStatusFailed  TransactionStatusEnum = "Failed"
)

func (x TransactionTypeEnum) String() string {
	return string(x)
}
