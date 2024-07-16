package sms

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/gin-gonic/gin"
	kafkaconnector "github.com/rosas99/monster/pkg/streams/connector/kafka"
	"net/http"
	"time"

	"github.com/rosas99/monster/pkg/log"
	genericoptions "github.com/rosas99/monster/pkg/options"
	"google.golang.org/grpc"
)

type Server interface {
	RunOrDie()
	GracefulStop()
}

type HTTPServer struct {
	srv         *http.Server
	httpOptions *genericoptions.HTTPOptions
	tlsOptions  *genericoptions.TLSOptions
}

type GRPCServer struct {
	srv  *grpc.Server
	opts *genericoptions.GRPCOptions
}

type MqServer struct {
	srv  *kafkaconnector.Queue
	opts *genericoptions.KafkaOptions
}

func NewHTTPServer(
	httpOptions *genericoptions.HTTPOptions,
	tlsOptions *genericoptions.TLSOptions,
	g *gin.Engine,
) (*HTTPServer, error) {

	httpsrv := &http.Server{Addr: httpOptions.Addr, Handler: g}
	var tlsConfig *tls.Config
	var err error
	if tlsOptions != nil && tlsOptions.UseTLS {
		tlsConfig, err = tlsOptions.TLSConfig()
		if err != nil {
			return nil, err
		}
		httpsrv.TLSConfig = tlsConfig
	}
	return &HTTPServer{srv: httpsrv, httpOptions: httpOptions, tlsOptions: tlsOptions}, nil
}

func (s *HTTPServer) RunOrDie() {
	log.Infof("Start to listening the incoming %s requests on %s", scheme(s.tlsOptions), s.httpOptions.Addr)
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalw("Failed to start http(s) server", "err", err)
	}
}

func (s *HTTPServer) GracefulStop() {
	// creates a context (ctx) for notifying the server goroutine,
	//which has 10 seconds to complete the current request being processed.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Errorw(err, "Failed to gracefully shutdown http(s) server")
	}
}

func NewMqServer(KafkaOptions *genericoptions.KafkaOptions, handler kafkaconnector.ConsumeHandler) (MqServer, error) {
	consumer, err := kafkaconnector.NewQueue(KafkaOptions, handler)
	if err != nil {
		return MqServer{}, err
	}

	return MqServer{srv: consumer, opts: KafkaOptions}, nil
}

func (s *MqServer) RunOrDie() {
	s.srv.Start()
}

func (s *MqServer) GracefulStop() {
	log.Infof("Gracefully stop mq server")
	s.srv.Stop()
}
