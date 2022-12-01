package constant

type XenditStatusEnum string

const (
	XenditStatusPending   XenditStatusEnum = "PENDING"
	XenditStatusVoided    XenditStatusEnum = "VOIDED"
	XenditStatusCompleted XenditStatusEnum = "COMPLETED"
	XenditStatusFailed    XenditStatusEnum = "FAILED"
	XenditStatusExpired   XenditStatusEnum = "EXPIRED"
)
