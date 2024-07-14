package service

import (
	"github.com/rosas99/monster/internal/sms/biz"
)

// SmsServerService is a struct that implements the v1.UnimplementedSmsServerServer interface
// and holds the business logic, represented by a IBiz instance.
type SmsServerService struct {
	biz biz.IBiz
}

// NewSmsServerService is a constructor function that takes a IBiz instance
// as an input and returns a new SmsServerService instance.
func NewSmsServerService(biz biz.IBiz) *SmsServerService {
	return &SmsServerService{biz: biz}
}
