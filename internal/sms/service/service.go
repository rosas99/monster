// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package service

import (
	"context"
	"github.com/rosas99/monster/internal/sms/biz"
	pb "github.com/rosas99/monster/pkg/api/sms/v1"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type SmsServerService struct {
	biz biz.IBiz
	v1.UnimplementedSmsServerServer
}

func (s *SmsServerService) DeleteOrder(ctx context.Context, request *pb.CreateTemplateRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SmsServerService) mustEmbedUnimplementedSmsServerServer() {
	//TODO implement me
	panic("implement me")
}

func NewSmsServerService(biz biz.IBiz) *SmsServerService {
	return &SmsServerService{biz: biz}
}
