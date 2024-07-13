package service

import (
	"context"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserCenterService) ChangePassword(ctx context.Context, rq *v1.ChangePasswordRequest) error {
	log.C(ctx).Infow("CreateOrder function called")
	return s.biz.Users().ChangePassword(ctx, rq)
}

func (s *UserCenterService) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginResponse, error) {
	return s.biz.Users().Login(ctx, rq)
}

func (s *UserCenterService) Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	log.C(ctx).Infow("GetOrder function called")
	return s.biz.Users().Create(ctx, rq)
}

func (s *UserCenterService) Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	log.C(ctx).Infow("GetOrder function called")
	return s.biz.Users().Get(ctx, rq)
}

func (s *UserCenterService) List(ctx context.Context, rq *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	log.C(ctx).Infow("UpdateOrder function called")
	return s.biz.Users().List(ctx, rq)
}

func (s *UserCenterService) Update(ctx context.Context, rq *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	log.C(ctx).Infow("UpdateOrder function called")
	return &emptypb.Empty{}, s.biz.Users().Update(ctx, rq)
}

func (s *UserCenterService) Delete(ctx context.Context, rq *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	log.C(ctx).Infow("DeleteOrder function called")
	return &emptypb.Empty{}, s.biz.Users().Delete(ctx, rq)
}
