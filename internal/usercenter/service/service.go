package service

import (
	"github.com/rosas99/monster/internal/usercenter/biz"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

type UserCenterService struct {
	v1.UnimplementedUserCenterServer
	biz biz.IBiz
}

func NewUserCenterService(biz biz.IBiz) *UserCenterService {
	return &UserCenterService{biz: biz}
}
