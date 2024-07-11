package redis

import (
	"fmt"
)

const (
	MobileCount                  = "MOBILE_COUNT"
	TemplateCount                = "TEMPLATE_COUNT"
	TimeInterval                 = "TIME_INTERVAL"
	TemplateTypeVerificationCode = "TEMPLATE_TYPE_VERIFICATION_CODE"
	DELIMITER                    = ":"
)

// WrapperMobileCount  is used to build the key name in Redis.
func WrapperMobileCount(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", MobileCount, DELIMITER, templateCode, DELIMITER, mobile)
}

// WrapperTemplateCount  is used to build the key name in Redis.
func WrapperTemplateCount(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", TemplateCount, DELIMITER, templateCode, DELIMITER, mobile)
}

// WrapperTimeInterval  is used to build the key name in Redis.
func WrapperTimeInterval(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", TimeInterval, DELIMITER, templateCode, DELIMITER, mobile)
}

// WrapperTimeInterval  is used to build the key name in Redis.
func WrapperCode(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", TemplateTypeVerificationCode, DELIMITER, templateCode, DELIMITER, mobile)
}

const (
	TemplateM = "TEMPLATE_M"
)

// WrapperTemplateM  is used to build the key name in Redis.
func WrapperTemplateM(templateCode string) string {
	return fmt.Sprintf("%s%s%s", TemplateM, DELIMITER, templateCode)
}
