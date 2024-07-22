package usercenter

import (
	"context"
	"github.com/rosas99/monster/pkg/log"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"

	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

var (
	once sync.Once
	cli  *impl
)

type Gender impl

// Interface is an interface that presents a subset of the usercenter API.
type Interface interface {
	Auth(ctx context.Context, token string) (userID string, err error)
}

// impl is an implementation of Interface.
type impl struct {
	client v1.UserCenterClient
}

type Impl = impl

var (
	addr  = flag.String("addr", "127.0.0.1:3380", "The address to connect to.")
	limit = flag.Int64("limit", 10, "Limit to list users.")
)

// NewUserCenterServer creates a new client to work with sms services.
func NewUserCenterServer() *impl {
	flag.Parse()
	once.Do(func() {
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalw("Did not connect", "err", err)
		}
		defer conn.Close()

		client := v1.NewUserCenterClient(conn)
		cli = &impl{client}
	})

	return cli
}

// GetClient returns the globally initialized client.
func GetClient() *impl {
	return cli
}

func (i *impl) Auth(ctx context.Context, token string) (userID string, err error) {
	rq := &v1.AuthzRequest{Token: token}
	resp, err := i.client.Authorize(ctx, rq)
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
