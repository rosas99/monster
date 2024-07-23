package types

import "time"

const (
	MessageCountForTemplatePerDay = "MESSAGE_COUNT_FOR_TEMPLATE_PER_DAY"
	MessageCountForMobilePerDay   = "MESSAGE_COUNT_FOR_MOBILE_PER_DAY"
	TimeIntervalForMobilePerDay   = "TIME_INTERVAL_FOR_MOBILE_PER_DAY"
)

// ProviderType defines an enumerated type for different cloud service providers.
type ProviderType string

// defines a group of constants for cloud service providers.

const (
	ProviderTypeALIYUN ProviderType = "AILIYUN"

	ProviderTypeDummy ProviderType = "DUMMY"
)

const (
	ErrorStatus = "ERROR"
)

const (
	LimitLeftTime = time.Hour * 24
)

// TemplateMsgRequest defines a template message request for kafka queue.
type TemplateMsgRequest struct {
	PhoneNumber  string   `json:"phoneNumber"`
	SendTime     string   `json:"sendTime"`
	Content      string   `json:"content"`
	SignName     string   `json:"signName"`
	DestCode     string   `json:"destCode"`
	SequenceId   int64    `json:"sequenceId"`
	RequestId    string   `json:"requestId"`
	TemplateCode string   `json:"templateCode"`
	Providers    []string `json:"providers"`
}

// UplinkMsgRequest defines an uplink message request for kafka queue.
type UplinkMsgRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	SendTime    string `json:"sendTime"`
	Content     string `json:"content"`
	SignName    string `json:"signName"`
	DestCode    string `json:"destCode"`
	SequenceId  int64  `json:"sequenceId"`
	RequestId   string `json:"requestId"`
}
