package service

import (
	"context"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/log"
)

func (s *UserCenterService) Authorize(ctx context.Context, rq *v1.LoginRequest2) (*v1.LoginResponse2, error) {
	log.C(ctx).Infow("CreateOrder function called")
	return s.biz.Users().Authorize(ctx, rq)
	//return &v1.LoginResponse{Token: "scs"}, nil
}
