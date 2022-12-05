package constant

type XenditStatusEnum string

const (
	XenditStatusPending   XenditStatusEnum = "PENDING"
	XenditStatusPaid      XenditStatusEnum = "PAID"
	XenditStatusSettled   XenditStatusEnum = "SETTLED"
	XenditStatusExpired   XenditStatusEnum = "EXPIRED"
	XenditStatusFailed    XenditStatusEnum = "FAILED"
	XenditStatusCompleted XenditStatusEnum = "COMPLETED"
)

func (x XenditStatusEnum) String() string {
	return string(x)
}
