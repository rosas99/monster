package app

import (
	"github.com/rosas99/monster/internal/usercenter"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/rosas99/monster/cmd/monster-usercenter/app/options"
	"github.com/rosas99/monster/pkg/app"
)

const commandDesc = `The sms server is a standard, specification-compliant demo 
example of the onex service.

Find more onex-sms information at:
    https://"github.com/rosas99/monster/blob/master/docs/guide/en-US/cmd/onex-sms.md`

// NewApp creates an App object with default parameters.
func NewApp(name string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(name, "Launch a onex sms server",
		app.WithDescription(commandDesc),
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func() error {
		cfg, err := opts.Config()
		if err != nil {
			return err
		}

		return Run(cfg, genericapiserver.SetupSignalHandler())
	}
}

// Run runs the specified APIServer. This should never exit.
func Run(c *usercenter.Config, stopCh <-chan struct{}) error {
	server, err := c.Complete().New()
	if err != nil {
		return err
	}

	return server.Run(stopCh)
}
