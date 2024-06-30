// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

// gateway is the backend portal of onex.
package main

import (
	_ "k8s.io/component-base/metrics/prometheus/clientgo" // load all the prometheus client-go plugins
	_ "k8s.io/component-base/metrics/prometheus/version"  // for version metric registration

	"github.com/rosas99/monster/cmd/monster-pump/app"
)

func main() {
	app.NewApp("monster-pump").Run()
}
