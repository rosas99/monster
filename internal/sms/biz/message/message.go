// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package message

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/checker"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/store"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/id"
	"github.com/rosas99/monster/pkg/log"
	"time"
)

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

// todo 生成请求
// Create 是 OrderBiz 接口中 `Create` 方法的实现.
func (b *messageBiz) Send(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {
	var templateMsgRequest types.TemplateMsgRequest
	templateMsgRequest.RequestId = b.idt.Token(ctx)
	// todo 先使用redis保存 后续再考虑使用本地缓存
	// todo 参考cache服务如何实现
	// 从本地缓存查询模板
	l := b.logger
	m := model.HistoryM{}

	templateM, err := b.ds.Templates().Get(ctx, rq.TemplateCode)
	if err != nil {
		l.LogHistory(&m)
	}

	// 如果是验证码，缓存验证码
	if templateM.Type == "VERIFICATION" {
		// todo 生成6位随机验证码
		rq.Code = id.RandomNumeric(6)

	}

	// 组装短信
	_ = copier.Copy(&templateMsgRequest, rq)
	//templateMsgRequest.RequestId = strconv.FormatUint(snow.GenerateId(), 10)

	// 从本地缓存获取限流配置 // 忽略count返回
	_, cfgList, err := b.ds.Configurations().List(ctx, rq.TemplateCode)
	if err != nil {
	}

	// 规则校验
	err = b.rule.CheckRules(templateM, rq.Mobile, cfgList)
	// 这里只使用err
	if err != nil {
		l.LogHistory(&m)
	}

	// 根据类型发送短信到对应mq
	// 异步记录短信历史
	key := factory.WrapperCode(rq.TemplateCode, rq.Mobile)
	b.rds.Set(ctx, key, rq.Code, time.Hour*24)

	l.WriteSendMessage(ctx, &templateMsgRequest)

	log.C(ctx).Infof("test")

	message := map[string]any{
		"test":  "value1",
		"other": 123,
	}
	l.LogKpi(message)
	// todo log记录短信延时
	return &v1.CreateTemplateResponse{OrderID: templateM.ID}, nil
}

func (b *messageBiz) CodeVerify(ctx context.Context, rq *v1.VerifyCodeRequest) (*v1.CommonResponse, error) {

	key := factory.WrapperCode(rq.Mobile, rq.TemplateCode)
	code, err := b.rds.Get(ctx, key).Result()
	if rq.Code != code {
		return &v1.CommonResponse{Code: 500, Msg: "fail"}, err
	}
	b.rds.Del(ctx, key)
	message := map[string]any{
		"test":  "value1",
		"other": 123,
	}
	b.logger.LogKpi(message)
	return &v1.CommonResponse{Code: 0, Msg: "success"}, nil

}

// 报告可以慢慢接收
func (b *messageBiz) AILIYUNReport(ctx context.Context, rq *v1.AILIYUNReportListRequest) (*v1.CommonResponse, error) {
	for _, item := range rq.AILIYUNReportList {
		filter := make(map[string]any)
		filter["message_id"] = item.BizId
		count, list, _ := b.ds.Histories().List(ctx, "", meta.WithFilter(filter))
		if count > 0 {
			history := list[0]
			marshal, _ := json.Marshal(history)
			history.Report = string(marshal)
		}
	}
	return &v1.CommonResponse{Code: 0, Msg: "success"}, nil
}
