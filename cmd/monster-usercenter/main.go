// fakeserver is a standard, specification-compliant demo example of the onex service.
// fakeserver is also a gRPC and HTTP server.
package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/rosas99/monster/cmd/monster-usercenter/app"
)

func main() {
	app.NewApp("monster-usercenter").Run()
}
