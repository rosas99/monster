// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package usercenter

import (
	"context"
	"github.com/rosas99/monster/pkg/log"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc/credentials/insecure"
	"sync"

	v1 "github.com/rosas99/monster/pkg/api/fakeserver/v1"
	"google.golang.org/grpc"
)

var (
	once sync.Once
	cli  *impl
)

type Gender impl

// Interface is an interface that presents a subset of the usercenter API.
type Interface interface {
	Auth(ctx context.Context, token string, obj, act string) (string, bool, error)
}

// impl is an implementation of Interface.
type impl struct {
	client v1.FakeServerClient
}

type Impl = impl

var (
	addr  = flag.String("addr", "localhost:9090", "The address to connect to.")
	limit = flag.Int64("limit", 10, "Limit to list users.")
)

// NewUserCenter creates a new client to work with usercenter services.
func NewFakeServer() *impl {
	flag.Parse()
	once.Do(func() {
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalw("Did not connect", "err", err)
		}
		defer conn.Close()

		rpcclient := v1.NewFakeServerClient(conn)
		cli = &impl{rpcclient}
	})

	return cli
}

// GetClient returns the globally initialized client.
func GetClient() *impl {
	return cli
}

func (i *impl) Authenticate(ctx context.Context, token string) (userID string, err error) {
	rq := &v1.DeleteOrderRequest{}
	resp, err := i.client.DeleteOrder(ctx, rq)
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
