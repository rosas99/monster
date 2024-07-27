package usercenter

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/pkg/known"
	"github.com/rosas99/monster/internal/pkg/middleware/header"
	"github.com/rosas99/monster/internal/pkg/middleware/trace"
	"github.com/rosas99/monster/internal/usercenter/biz"
	"github.com/rosas99/monster/pkg/auth"
	"github.com/rosas99/monster/pkg/token"
	//"github.com/rosas99/monster/internal/usercenter/middleware/validate"
	"github.com/rosas99/monster/internal/usercenter/service"
	"github.com/rosas99/monster/internal/usercenter/store"
	"github.com/rosas99/monster/internal/usercenter/store/mysql"
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
	JwtSecret    string
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
	config  completedConfig
}

// New returns a new instance of Server from the given config.
// New initializes and returns a new SmsServer instance.
func (c completedConfig) New() (*SmsServer, error) {
	var ds store.IStore

	// Copy MySQL options from the configuration.
	var dbOptions db.MySQLOptions
	_ = copier.Copy(&dbOptions, c.MySQLOptions)

	// Initialize the MySQL database instance.
	ins, err := db.NewMySQL(&dbOptions)
	if err != nil {
		return nil, err
	}
	ds = mysql.NewStore(ins)

	// Copy Redis options from the configuration.
	var redisOptions db.RedisOptions
	value := &c.Config.RedisOptions
	_ = copier.Copy(&redisOptions, value)

	// Initialize the Redis database instance.
	rds, err := db.NewRedis(&redisOptions)
	if err != nil {
		return nil, err
	}

	// Initialize the token package with the JWT secret key.
	token.Init(c.JwtSecret, known.XUsernameKey)

	// Create a business layer instance using the initialized data stores.
	biz := biz.NewBiz(ds, rds)
	authz, err := auth.NewAuthz(ins)
	srv := service.NewUserCenterService(biz, authz)

	// Set the Gin mode to release for production.
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine.
	g := gin.New()

	// Initialize routes for the Gin engine.
	installRouters(g, srv)

	// Initialize the HTTP server.
	httpsrv, err := NewHTTPServer(c.HTTPOptions, c.TLSOptions, g)
	if err != nil {
		return nil, err
	}

	// Initialize the gRPC server.
	grpcsrv, err := NewGRPCServer(c.GRPCOptions, c.TLSOptions, srv)
	if err != nil {
		return nil, err
	}

	// Define middleware for the Gin engine including recovery, no-cache, CORS, security, and trace ID.
	mws := []gin.HandlerFunc{gin.Recovery(), header.NoCache, header.Cors, header.Secure, trace.TraceID()}
	g.Use(mws...)

	// Start the gRPC server before the HTTP server since the HTTP server might depend on it.
	go grpcsrv.RunOrDie()

	// Return the newly created SmsServer instance.
	return &SmsServer{grpcsrv: grpcsrv, httpsrv: httpsrv, config: c}, nil
}

func (s *SmsServer) Run(stopCh <-chan struct{}) error {

	log.Infof("Successfully start pump server")

	go s.httpsrv.RunOrDie()

	<-stopCh

	// The most gracefully way is to shutdown the dependent service first,
	// and then shutdown the depended service.
	s.httpsrv.GracefulStop()
	s.grpcsrv.GracefulStop()

	return nil
}
