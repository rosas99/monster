package service

import (
	"context"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ChangePassword changes the password for a user by calling the corresponding method in the business logic layer.
func (s *UserCenterService) ChangePassword(ctx context.Context, username string, rq *v1.ChangePasswordRequest) error {
	log.C(ctx).Infow("ChangePassword function called")
	return s.biz.Users().ChangePassword(ctx, username, rq)
}

// Login authenticates a user and returns a login response.
func (s *UserCenterService) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginResponse, error) {
	return s.biz.Users().Login(ctx, rq)
}

// Create creates a new user and returns a create user response.
func (s *UserCenterService) Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	log.C(ctx).Infow("CreateUser function called")
	return s.biz.Users().Create(ctx, rq)
}

// Get retrieves a user's information and returns a get user response.
func (s *UserCenterService) Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	log.C(ctx).Infow("GetUser function called")
	return s.biz.Users().Get(ctx, rq)
}

// List lists users and returns a list user response.
func (s *UserCenterService) List(ctx context.Context, rq *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	log.C(ctx).Infow("ListUsers function called")
	return s.biz.Users().List(ctx, rq)
}

// Update updates a user's information and returns an empty response.
func (s *UserCenterService) Update(ctx context.Context, rq *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	log.C(ctx).Infow("UpdateUser function called")
	return &emptypb.Empty{}, s.biz.Users().Update(ctx, rq)
}

// Delete deletes a user and returns an empty response.
func (s *UserCenterService) Delete(ctx context.Context, rq *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	log.C(ctx).Infow("DeleteUser function called")
	return &emptypb.Empty{}, s.biz.Users().Delete(ctx, rq)
}
