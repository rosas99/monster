// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package service

import (
	"github.com/rosas99/monster/internal/sms/biz"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

type SmsServerService struct {
	biz biz.IBiz
	v1.UnimplementedSmsServerServer
}

func (s *SmsServerService) mustEmbedUnimplementedSmsServerServer() {
	//TODO implement me
	panic("implement me")
}

func NewSmsServerService(biz biz.IBiz) *SmsServerService {
	return &SmsServerService{biz: biz}
}
