// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

//go:build wireinject
// +build wireinject

package nightwatch

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/google/wire"

	gwstore "github.com/rosas99/monster/internal/gateway/store"
	"github.com/rosas99/monster/internal/pkg/client/store"
	ucstore "github.com/rosas99/monster/internal/usercenter/store"
	"github.com/rosas99/monster/pkg/db"
)

func wireStoreClient(*db.MySQLOptions) (store.Interface, error) {
	wire.Build(
		db.ProviderSet,
		store.ProviderSet,
		gwstore.ProviderSet,
		ucstore.ProviderSet,
	)

	return nil, nil
}
