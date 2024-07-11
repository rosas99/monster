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

// SmsServerService is a struct that implements the v1.UnimplementedSmsServerServer interface
// and holds the business logic, represented by a IBiz instance.
type SmsServerService struct {
	biz biz.IBiz
	v1.UnimplementedSmsServerServer
}

// NewSmsServerService is a constructor function that takes a IBiz instance
// as an input and returns a new SmsServerService instance.
func NewSmsServerService(biz biz.IBiz) *SmsServerService {
	return &SmsServerService{biz: biz}
}
