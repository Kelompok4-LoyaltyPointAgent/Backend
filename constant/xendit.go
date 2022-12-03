package constant

type XenditStatusEnum string

const (
	XenditStatusPending XenditStatusEnum = "PENDING"
	XenditStatusPaid    XenditStatusEnum = "PAID"
	XenditStatusSettled XenditStatusEnum = "SETTLED"
	XenditStatusExpired XenditStatusEnum = "EXPIRED"
)

func (x XenditStatusEnum) String() string {
	return string(x)
}
