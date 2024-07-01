// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package biz

//go:generate mockgen -destination mock_biz.go -package biz github.com/rosas99/monster/internal/fakeserver/biz IBiz

import (
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/sms/biz/template"
	"github.com/rosas99/monster/internal/usercenter/biz/user"
	"github.com/rosas99/monster/internal/usercenter/store"
	"github.com/segmentio/kafka-go"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	Templates() template.TemplateBiz
}

// biz 是 IBiz 的一个具体实现.
type Biz struct {
	ds          store.IStore
	rds         *redis.Client
	kafkaWriter *kafka.Writer
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*Biz)(nil)

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore, rds *redis.Client) *Biz {
	return &Biz{ds: ds, rds: rds}
}

// Orders 返回一个实现了 OrderBiz 接口的实例.
func (b *Biz) Users() template.TemplateBiz {
	return user.New(b.ds, b.rds)
}
