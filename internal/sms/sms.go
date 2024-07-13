package sms

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/pkg/client/usercenter"
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
	ailiyunoptions "github.com/rosas99/monster/pkg/sdk/ailiyun"
)

// Config represents the configuration of the service.
type Config struct {
	GRPCOptions       *genericoptions.GRPCOptions
	HTTPOptions       *genericoptions.HTTPOptions
	TLSOptions        *genericoptions.TLSOptions
	MySQLOptions      *genericoptions.MySQLOptions
	RedisOptions      *genericoptions.RedisOptions
	KafkaOptions      *genericoptions.KafkaOptions
	Address           string
	Accounts          map[string]string
	AiliyunUrl        string
	AiliyunSmsOptions *ailiyunoptions.SmsOptions
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

// New returns a new instance of SmsServer from the given config.
func (c completedConfig) New() (*SmsServer, error) {
	var ds store.IStore

	var dbOptions db.MySQLOptions
	_ = copier.Copy(&dbOptions, c.MySQLOptions)
	ins, err := db.NewMySQL(&dbOptions)
	if err != nil {
		return nil, err
	}
	ds = mysql.NewStore(ins)

	var redisOptions db.RedisOptions
	value := &c.Config.RedisOptions
	_ = copier.Copy(&redisOptions, value)
	rds, err := db.NewRedis(&redisOptions)
	if err != nil {
		return nil, err
	}

	// registers message check rules
	factory := checker.NewRuleFactory()
	factory.RegisterRule(types.MessageCountForTemplatePerDay, checker.NewMessageCountForTemplateRule(ds, rds))
	factory.RegisterRule(types.MessageCountForMobilePerDay, checker.NewMessageCountForMobileRule(ds, rds))
	factory.RegisterRule(types.TimeIntervalForMobilePerDay, checker.NewTimeIntervalForMobileRule(ds, rds))

	// creates  a logger instance
	// todo 其他options
	l, err := logger.NewLogger(c.KafkaOptions, c.KafkaOptions, ds.Histories())
	if err != nil {
		return nil, err
	}

	// registers sms providers
	provider := providerFactory.NewProviderFactory()
	provider.RegisterProvider(types.ProviderTypeALIYUN, providerFactory.NewAILIYUNProvider(rds, l, c.AiliyunSmsOptions))

	// creates an idempotent instance
	idt, err := idempotent.NewIdempotent(rds)
	if err != nil {
		return nil, err
	}

	bizIns := biz.NewBiz(ds, rds, idt, l)
	srv := service.NewSmsServerService(bizIns)

	// Sets the running mode for the Gin
	gin.SetMode(gin.ReleaseMode)
	// create a gin engine
	g := gin.New()

	usercenter.NewUserCenterServer()

	installRouters(g, srv)
	mws := []gin.HandlerFunc{
		gin.Recovery(), header.NoCache, header.Cors, header.Secure,
		trace.TraceID(), nil, validate.Validation(ds),
	}
	// add gin middlewares
	g.Use(mws...)

	httpsrv, err := NewHTTPServer(c.HTTPOptions, c.TLSOptions, g)
	if err != nil {
		return nil, err
	}

	grpcsrv, err := NewGRPCServer(c.GRPCOptions, c.TLSOptions, srv)
	if err != nil {
		return nil, err
	}

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

// Run is a method of the SmsServer struct that starts the server.
func (s *SmsServer) Run(stopCh <-chan struct{}) error {

	log.Infof("Successfully start sms server")
	go s.httpsrv.RunOrDie()

	<-stopCh

	log.Infof("Gracefully shutting down sms server ...")

	// The most gracefully way is to shut down the dependent service first,
	// and then shutdown the depended on service.
	s.httpsrv.GracefulStop()
	s.grpcsrv.GracefulStop()
	s.mqsrv.GracefulStop()
	s.mqsrv2.GracefulStop()

	return nil
}
