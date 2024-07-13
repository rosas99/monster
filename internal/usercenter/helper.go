package usercenter

import (
	genericoptions "github.com/rosas99/monster/pkg/options"
)

func scheme(opts *genericoptions.TLSOptions) string {
	scheme := "http"
	if opts != nil && opts.UseTLS {
		scheme = "https"
	}

	return scheme
}
