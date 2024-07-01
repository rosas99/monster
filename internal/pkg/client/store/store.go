// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package store

import (
	ucstore "github.com/rosas99/monster/internal/usercenter/store"
	"sync"
)

var (
	once sync.Once
	S    *datastore
)

// Interface defines the storage interface.
type Interface interface {
	UserCenter() ucstore.IStore
	// 导入了store 以进行查询
}

type datastore struct {
	uc ucstore.IStore
}

var _ Interface = (*datastore)(nil)

func (ds *datastore) UserCenter() ucstore.IStore {
	return ds.uc
}

func NewStore(uc ucstore.IStore) *datastore {
	once.Do(func() {
		S = &datastore{uc: uc}
	})

	return S
}
