// fakeserver is a standard, specification-compliant demo example of the onex service.
// fakeserver is also a gRPC and HTTP server.
package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/rosas99/monster/cmd/monster-nightwatch/app"
)

func main() {
	app.NewApp("monster-nightwatch").Run()
}
