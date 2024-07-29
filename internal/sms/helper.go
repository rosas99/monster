package sms

import (
	"fmt"
	genericoptions "github.com/rosas99/monster/pkg/options"
)

func scheme(opts *genericoptions.TLSOptions) string {
	scheme := "http"
	if opts != nil && opts.UseTLS {
		scheme = "https"
	}

	return scheme
}

const (
	MobileCount                  = "MOBILE_COUNT"
	TemplateCount                = "TEMPLATE_COUNT"
	TimeInterval                 = "TIME_INTERVAL"
	TemplateTypeVerificationCode = "TEMPLATE_TYPE_VERIFICATION_CODE"
	DELIMITER                    = ":"

	TemplateM   = "TEMPLATE_M"
	TemplateCfg = "TEMPLATE_CONFIGURATION"
)

// WrapperTemplateCount  is used to build the key name in Redis.
func WrapperTemplateCount(templateCode string) string {
	return fmt.Sprintf("%s%s%s", MobileCount, DELIMITER, templateCode)
}

// WrapperMobileCount  is used to build the key name in Redis.
func WrapperMobileCount(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", TemplateCount, DELIMITER, templateCode, DELIMITER, mobile)
}

// WrapperTimeInterval  is used to build the key name in Redis.
func WrapperTimeInterval(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", TimeInterval, DELIMITER, templateCode, DELIMITER, mobile)
}

// WrapperCode  is used to build the key name in Redis.
func WrapperCode(templateCode, mobile string) string {
	return fmt.Sprintf("%s%s%s%s%s", TemplateTypeVerificationCode, DELIMITER, templateCode, DELIMITER, mobile)
}

// WrapperTemplate  is used to build the key name in Redis.
func WrapperTemplate(templateCode string) string {
	return fmt.Sprintf("%s%s%s", TemplateM, DELIMITER, templateCode)
}

func WrapperTemplateCfg(templateCode string) string {
	return fmt.Sprintf("%s%s%s", TemplateCfg, DELIMITER, templateCode)
}
