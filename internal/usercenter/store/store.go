// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package store

//go:generate mockgen -destination mock_store.go -package store github.com/rosas99/monster/internal/fakeserver/store IStore,OrderStore

import (
	"context"
	"sync"

	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/usercenter/model"
)

// IStore 定义了 Store 层需要实现的方法.
type IStore interface {
	Users() UserStore
}

// OrderStore 定义了 order 模块在 store 层所实现的方法.

type UserStore interface {
	Create(ctx context.Context, order *model.UserM) error
	Get(ctx context.Context, username string) (*model.UserM, error)
	Update(ctx context.Context, user *model.UserM) error
	List(ctx context.Context, opts ...meta.ListOption) (count int64, ret []*model.UserM, err error)
	Delete(ctx context.Context, id int64) error
}

var (
	once sync.Once
	// 全局变量，方便其它包直接调用已初始化好的 S 实例.
	S IStore
)

// SetStore set the onex-fakeserver store instance in a global variable `S`.
// Direct use the global `S` is not recommended as this may make dependencies and calls unclear.
func SetStore(store IStore) {
	once.Do(func() {
		S = store
	})
}
