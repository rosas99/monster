// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package message

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/rule"
	"github.com/rosas99/monster/internal/sms/store"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/id"
	"github.com/rosas99/monster/pkg/log"
)

type MessageBiz interface {
	Send(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error)
}

// OrderBiz 接口的实现.
type messageBiz struct {
	ds store.IStore
	// todo writer
	logger *Logger
	rds    *redis.Client
	rule   *rule.RuleFactory
	idt    *idempotent.Idempotent
}

// 确保 orderBiz 实现了 OrderBiz 接口.
var _ MessageBiz = (*messageBiz)(nil)

// New 创建一个实现了 OrderBiz 接口的实例.
func New(ds store.IStore, logger *Logger, rds *redis.Client, idt *idempotent.Idempotent) *messageBiz {
	return &messageBiz{ds: ds, logger: logger, rds: rds}
}

type TemplateMsgRequest struct {
	Matcher   string     `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	Request   []any      `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Result    bool       `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	Explains  [][]string `protobuf:"bytes,4,opt,name=explains,proto3" json:"explains,omitempty"`
	Timestamp int64      `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	RequestId string     `protobuf:"bytes,6,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

// todo 生成请求
// Create 是 OrderBiz 接口中 `Create` 方法的实现.
func (b *messageBiz) Send(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {
	// todo request只在controller层用，biz需要用到什么参数单独取出来
	var templateMsgRequest TemplateMsgRequest
	templateMsgRequest.RequestId = b.idt.Token(ctx)
	// todo 先使用redis保存 后续再考虑使用本地缓存
	// 从本地缓存查询模板
	templateM, err := b.ds.Templates().Get(ctx, rq.TemplateCode)
	if err != nil {

	}
	// 判断是否是校验码类型，如果校验码为空，则生成
	// 中间件生成请求id 保存到缓存池 mq消费时进行检查【使用本地线程变量】消费后清除

	// 雪花id生成短信id
	id := "testid"
	// 组装短信
	_ = copier.Copy(&templateMsgRequest, rq)
	// todo map参数转string
	templateMsgRequest.RequestId = id

	// 从本地缓存获取限流配置 // 忽略count返回
	_, cfgList, err := b.ds.Configurations().List(ctx, rq.TemplateCode)
	if err != nil {
	}

	// 规则校验
	err = b.rule.CheckRules(templateM, "mobiletest", cfgList)
	// 这里只使用err
	if err == nil {
		// 如果是验证码，缓存验证码
		// 根据类型发送短信到对应mq
		// 异步记录短信历史
	}

	// 记录失败日志
	log.C(ctx).Infof("test")
	// 保存失败历史
	b.logger.LogHistory("")

	return nil, err
	// todo log记录短信延时
	//return &v1.CreateTemplateResponse{OrderID: templateM.ID}, nil
}

func main() {
	// 整个可以设置为全局变量 只初始化一次
	options := func(*id.SonyflakeOptions) {
		id.WithSonyflakeMachineId(1) // 自定义机器ID，默认为自动检测
	}

	snowIns := id.NewSonyflake(options)
	id := snowIns.Id(context.Background())
	fmt.Print("id is :", id)
}
