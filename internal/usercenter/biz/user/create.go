package user

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/usercenter/model"
	v1 "github.com/rosas99/monster/pkg/api/usercenter/v2"
	"regexp"
)

func (b *userBiz) Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {

	var userM model.UserM
	_ = copier.Copy(&userM, rq)
	if err := b.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return nil, errors.New("old password is invalid")
		}
		return nil, nil
	}

	return &v1.CreateUserResponse{UserID: userM.ID}, nil
}
