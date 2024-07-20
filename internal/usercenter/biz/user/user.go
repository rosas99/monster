package user

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/usercenter/store"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

type IBiz interface {
	Authorize(ctx context.Context, rq *v1.AuthzRequest) (*v1.AuthzResponse, error)
	ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
	Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.CreateUserResponse, error)
	Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.GetUserResponse, error)
	List(ctx context.Context, rq *v1.ListUserRequest) (*v1.ListUserResponse, error)
	Update(ctx context.Context, rq *v1.UpdateUserRequest) error
	Delete(ctx context.Context, rq *v1.DeleteUserRequest) error
}
type userBiz struct {
	ds  store.IStore
	rds *redis.Client
}

func New(ds store.IStore, rds *redis.Client) *userBiz {
	return &userBiz{ds: ds, rds: rds}
}
