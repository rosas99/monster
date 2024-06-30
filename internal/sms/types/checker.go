package types

import "time"

const (
	LimitLeftTime = time.Hour * 24
)
const (
	MessageCountForTemplatePerDay = "MESSAGE_COUNT_FOR_TEMPLATE_PER_DAY"
	MessageCountForMobilePerDay   = "MESSAGE_COUNT_FOR_MOBILE_PER_DAY"
	TimeIntervalForMobilePerDay   = "TIME_INTERVAL_FOR_MOBILE_PER_DAY"
)

// Response  模拟验证失败的原因
type Response struct {
	resultCode string
	errorMsg   string
}

// Request  模拟验证请求
type Request struct {
	mobile       string
	id           int64
	templateCode string
	limitValue   int64
}
