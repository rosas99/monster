package service

import (
	"github.com/rosas99/monster/internal/usercenter/biz"
)

type UserCenterService struct {
	biz biz.IBiz
}

func NewUserCenterService(biz biz.IBiz) *UserCenterService {
	return &UserCenterService{biz: biz}
}
