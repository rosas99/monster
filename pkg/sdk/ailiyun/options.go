package ailiyun

import (
	"github.com/spf13/pflag"
)

// Options contains configuration options for logging.
type Options struct {
	Url string `json:"format,omitempty" mapstructure:"format"`
}

// NewOptions creates a new Options object with default values.
func NewOptions() *Options {
	return &Options{
		Url: "console",
	}
}

// Validate verifies flags passed to LogsOptions.
func (o *Options) Validate() []error {
	var errs []error

	return errs
}

// AddFlags adds command line flags for the configuration.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Url, "log.format", o.Url, "Log output `FORMAT`, support plain or json format.")
}
