package service

import (
	"context"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/log"
)

func (s *UserCenterService) Authorize(ctx context.Context, rq *v1.AuthzRequest) (*v1.AuthzResponse, error) {
	log.C(ctx).Infow("CreateOrder function called")
	return s.biz.Users().Authorize(ctx, rq)
}
