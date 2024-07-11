package types

import "time"

const (
	LimitLeftTime = time.Hour * 24
)

// defines a group of constants for message configuration.

const (
	MessageCountForTemplatePerDay = "MESSAGE_COUNT_FOR_TEMPLATE_PER_DAY"
	MessageCountForMobilePerDay   = "MESSAGE_COUNT_FOR_MOBILE_PER_DAY"
	TimeIntervalForMobilePerDay   = "TIME_INTERVAL_FOR_MOBILE_PER_DAY"
)

const (
	CommonMessage       = "COMMON_MESSAGE"
	VerificationMessage = "VERIFICATION_MESSAGE"
)

// Request  模拟验证请求
type Request struct {
	Mobile       string
	Id           int64
	TemplateCode string
	LimitValue   int64
}
