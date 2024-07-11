// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package message

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/checker"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/store"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// MessageBiz
type MessageBiz interface {
	Send(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error)
	CodeVerify(ctx context.Context, rq *v1.VerifyCodeRequest) (*v1.CommonResponse, error)
	AILIYUNReport(ctx context.Context, rq *v1.AILIYUNReportListRequest) (*v1.CommonResponse, error)
}

// OrderBiz 接口的实现.
type messageBiz struct {
	ds     store.IStore
	logger *logger.Logger
	rds    *redis.Client
	rule   *checker.RuleFactory
	idt    *idempotent.Idempotent
}

// 确保 orderBiz 实现了 OrderBiz 接口.
var _ MessageBiz = (*messageBiz)(nil)

// New 创建一个实现了 OrderBiz 接口的实例.
func New(ds store.IStore, logger *logger.Logger, rds *redis.Client, idt *idempotent.Idempotent) *messageBiz {
	return &messageBiz{ds: ds, logger: logger, rds: rds}
}
