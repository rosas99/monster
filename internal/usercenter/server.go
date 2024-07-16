package usercenter

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"time"

	pb "github.com/rosas99/monster/pkg/api/usercenter/v1"
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

func NewHTTPServer(
	httpOptions *genericoptions.HTTPOptions,
	tlsOptions *genericoptions.TLSOptions,
	g *gin.Engine,
) (*HTTPServer, error) {

	// 创建 HTTP Server 实例
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Errorw(err, "Failed to gracefully shutdown http(s) server")
	}
}

func NewGRPCServer(
	grpcOptions *genericoptions.GRPCOptions,
	tlsOptions *genericoptions.TLSOptions,
	srv pb.UserCenterServer,
) (*GRPCServer, error) {
	dialOptions := []grpc.ServerOption{}
	if tlsOptions != nil && tlsOptions.UseTLS {
		tlsConfig, err := tlsOptions.TLSConfig()
		if err != nil {
			return nil, err
		}

		// grpc middleware
		dialOptions = append(dialOptions, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}

	grpcsrv := grpc.NewServer(dialOptions...)
	pb.RegisterUserCenterServer(grpcsrv, srv)
	reflection.Register(grpcsrv)

	return &GRPCServer{srv: grpcsrv, opts: grpcOptions}, nil
}

func (s *GRPCServer) RunOrDie() {
	lis, err := net.Listen("tcp", s.opts.Addr)
	if err != nil {
		log.Fatalw("Failed to listen", "err", err)
	}

	log.Infow("Start to listening the incoming requests on grpc address", "addr", s.opts.Addr)
	if err := s.srv.Serve(lis); err != nil {
		log.Fatalw(err.Error())
	}
}

func (s *GRPCServer) GracefulStop() {
	log.Infof("Gracefully stop grpc server")
	s.srv.GracefulStop()
}
