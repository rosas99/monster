package ailiyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/pflag"
)

// SmsOptions contains configuration options for logging.
type SmsOptions struct {
	AccessKeyId     string `json:"accessKeyId,omitempty" mapstructure:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret,omitempty" mapstructure:"accessKeySecret"`
}

// NewSmsOptions creates a new SmsOptions object with default values.
func NewSmsOptions() *SmsOptions {
	return &SmsOptions{
		AccessKeyId:     "console",
		AccessKeySecret: "console",
	}
}

// Validate verifies flags passed to LogsOptions.
func (o *SmsOptions) Validate() []error {
	var errs []error

	return errs
}

// AddFlags adds command line flags for the configuration.
func (o *SmsOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.AccessKeyId, "log.format", o.AccessKeySecret, "Log output `FORMAT`, support plain or json format.")
}

func (o *SmsOptions) NewSmsClient() (_result *dysmsapi.Client, _err error) {
	config := &openapi.Config{}
	config.AccessKeyId = tea.String(o.AccessKeySecret)
	config.AccessKeySecret = tea.String(o.AccessKeyId)
	_result = &dysmsapi.Client{}
	_result, _err = dysmsapi.NewClient(config)
	return _result, _err
}
