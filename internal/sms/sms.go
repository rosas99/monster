// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package sms

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/pkg/middleware/header"
	"github.com/rosas99/monster/internal/pkg/middleware/trace"
	"github.com/rosas99/monster/internal/sms/biz"
	"github.com/rosas99/monster/internal/sms/checker"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/middleware/validate"
	mqs "github.com/rosas99/monster/internal/sms/mqs"
	providerFactory "github.com/rosas99/monster/internal/sms/provider"
	"github.com/rosas99/monster/internal/sms/service"
	"github.com/rosas99/monster/internal/sms/store"
	"github.com/rosas99/monster/internal/sms/store/mysql"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/db"
	"github.com/rosas99/monster/pkg/log"
	genericoptions "github.com/rosas99/monster/pkg/options"
)

// Config represents the configuration of the service.
type Config struct {
	GRPCOptions  *genericoptions.GRPCOptions
	HTTPOptions  *genericoptions.HTTPOptions
	TLSOptions   *genericoptions.TLSOptions
	MySQLOptions *genericoptions.MySQLOptions
	RedisOptions *genericoptions.RedisOptions
	KafkaOptions *genericoptions.KafkaOptions
	Address      string
	Accounts     map[string]string
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() completedConfig {
	return completedConfig{cfg}
}

type completedConfig struct {
	*Config
}

// SmsServer represents the fake server.
type SmsServer struct {
	httpsrv Server
	grpcsrv Server
	mqsrv   MqServer
	mqsrv2  MqServer
	config  completedConfig
}

// New returns a new instance of Server from the given config.
func (c completedConfig) New() (*SmsServer, error) {
	var ds store.IStore

	var dbOptions db.MySQLOptions
	_ = copier.Copy(&dbOptions, c.MySQLOptions)

	ins, err := db.NewMySQL(&dbOptions)
	if err != nil {
		return nil, err
	}
	// todo 这里需要指定model
	//ins.AutoMigrate(&model.OrderM{})
	ds = mysql.NewStore(ins)

	var redisOptions db.RedisOptions
	value := &c.Config.RedisOptions
	_ = copier.Copy(&redisOptions, value)
	rds, err := db.NewRedis(&redisOptions)
	if err != nil {
		return nil, err
	}

	// 注册rule
	// todo 修改为new
	factory := checker.NewRuleFactory()
	factory.RegisterRule(types.MessageCountForTemplatePerDay, checker.NewMessageCountForTemplateRule(ds, rds))
	factory.RegisterRule(types.MessageCountForMobilePerDay, checker.NewMessageCountForMobileRule(ds, rds))
	factory.RegisterRule(types.TimeIntervalForMobilePerDay, checker.NewTimeIntervalForMobileRule(ds, rds))

	provider := providerFactory.NewProviderFactory()
	provider.RegisterProvider(types.ProviderTypeWE, providerFactory.NewWEProvider(rds))

	// todo 其他消费者配置
	l, err := logger.NewLogger(c.KafkaOptions, ds.Histories())
	if err != nil {
		return nil, err
	}

	//这里初始化所有writer 然后注入biz
	idt, err := idempotent.NewIdempotent(rds)
	if err != nil {
		return nil, err
	}

	biz := biz.NewBiz(ds, rds, idt, l)

	srv := service.NewSmsServerService(biz)

	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)

	// 创建 Gin 引擎
	g := gin.New()

	// 并初始化路由
	// 这里注册不同的路由可以分开，如是否使用人认证中间件，分别在use 认证中间件前后

	installRouters(g, srv)
	// 考虑在这里install consumer

	httpsrv, err := NewHTTPServer(c.HTTPOptions, c.TLSOptions, g)
	if err != nil {
		return nil, err
	}

	grpcsrv, err := NewGRPCServer(c.GRPCOptions, c.TLSOptions, srv)
	if err != nil {
		return nil, err
	}

	// gin.Recovery() 中间件，用来捕获任何 panic，并恢复
	mws := []gin.HandlerFunc{gin.Recovery(), header.NoCache, header.Cors, header.Secure,
		// todo 这里传入rds ds
		// 注意验证链路的顺序
		trace.TraceID(), nil, validate.Validation(ds)}
	// 添加中间件
	g.Use(mws...)

	logic := mqs.NewMessageConsumer(context.Background(), idt, l, provider)
	mqsrv, err := NewMqServer(c.KafkaOptions, logic, true)
	if err != nil {
		return nil, err
	}
	go mqsrv.RunOrDie()

	// todo 其他kafka options
	logic2 := mqs.NewUplinkMessageConsumer(context.Background(), ds, idt, l)
	mqsrv2, err := NewMqServer(c.KafkaOptions, logic2, true)
	if err != nil {
		return nil, err
	}
	go mqsrv.RunOrDie()
	go mqsrv2.RunOrDie()

	// Need start grpc server first. http server depends on grpc sever.
	go grpcsrv.RunOrDie()
	return &SmsServer{grpcsrv: grpcsrv, httpsrv: httpsrv, mqsrv: mqsrv, mqsrv2: mqsrv2, config: c}, nil
}

func (s *SmsServer) Run(stopCh <-chan struct{}) error {

	log.Infof("Successfully start pump server")
	go s.httpsrv.RunOrDie()

	<-stopCh

	// The most gracefully way is to shutdown the dependent service first,
	// and then shutdown the depended service.
	s.httpsrv.GracefulStop()
	s.grpcsrv.GracefulStop()
	s.mqsrv.GracefulStop()
	s.mqsrv2.GracefulStop()

	return nil
}
