package types

import (
	"fmt"
	"time"
)

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
	ProviderTypeDummy  ProviderType = "DUMMY"
)

const (
	ErrorStatus = "ERROR"
)

const (
	LimitLeftTime = time.Hour * 24
)

const (
	VerificationMessage = "VERIFICATION_MESSAGE"
	CommonMessage       = "COMMON_MESSAGE"
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

// WrapperTemplateCount  is used to build the key name in Redis.
func WrapperTemplateCount(templateCode string) string {
	return fmt.Sprintf("%s%s%s", MobileCount, DELIMITER, templateCode)
}

// WrapperTemplate  is used to build the key name in Redis.
func WrapperTemplate(templateCode string) string {
	return fmt.Sprintf("%s%s%s", TemplateM, DELIMITER, templateCode)
}

const (
	TimeInterval = "TIME_INTERVAL"

	TemplateCfg = "TEMPLATE_CONFIGURATION"

	TemplateCount = "TEMPLATE_COUNT"

	TemplateTypeVerificationCode = "TEMPLATE_TYPE_VERIFICATION_CODE"
)

const (
	MobileCount = "MOBILE_COUNT"

	DELIMITER = ":"

	TemplateM = "TEMPLATE_M"
)

// WrapperMobileCount  is used to build the key name in Redis.
func WrapperMobileCount(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", TemplateCount, DELIMITER, templateCode, DELIMITER, mobile)
}

// WrapperTimeInterval  is used to build the key name in Redis.
func WrapperTimeInterval(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", TimeInterval, DELIMITER, templateCode, DELIMITER, mobile)
}

func WrapperTemplateCfg(templateCode string) string {
	return fmt.Sprintf("%s%s%s", TemplateCfg, DELIMITER, templateCode)
}

// WrapperCode  is used to build the key name in Redis.
func WrapperCode(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", TemplateTypeVerificationCode, DELIMITER, templateCode, DELIMITER, mobile)
}
