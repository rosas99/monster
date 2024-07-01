// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package mysql

import (
	"sync"

	"gorm.io/gorm"

	"github.com/rosas99/monster/internal/sms/store"
)

var (
	once sync.Once
	// 全局变量，保存已被初始化的 *Datastore 实例.
	s *Datastore
)

// Datastore 是 IStore 的一个具体实现.
type Datastore struct {
	db *gorm.DB
}

// 确保 Datastore 实现了 store.IStore 接口.
var _ store.IStore = (*Datastore)(nil)

// NewStore 创建一个 store.IStore 类型的实例.
func NewStore(db *gorm.DB) *Datastore {
	// 确保 s 只被初始化一次
	once.Do(func() {
		s = &Datastore{db}
	})

	return s
}

func (ds *Datastore) Templates() store.TemplateStore {
	return newTemplates(ds.db)
}

func (ds *Datastore) Configurations() store.ConfigurationStore {
	return newConfigurations(ds.db)
}

// todo history
