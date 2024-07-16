package user

import (
	"context"
	"errors"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/auth"
	"github.com/rosas99/monster/pkg/token"
)

func (b *userBiz) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	user, err := b.ds.Users().Get(ctx, r.Username)
	if err != nil {
		return nil, errors.New("old password is invalid")

	}

	if err := auth.Compare(user.Password, r.Password); err != nil {
		return nil, errors.New("old password is invalid")
	}

	t, err := token.Sign(r.Username)
	if err != nil {
		return nil, errors.New("old password is invalid")
	}

	return &v1.LoginResponse{Token: t}, nil
}
