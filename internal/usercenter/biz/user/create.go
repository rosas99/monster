package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/pkg/errno"
	"github.com/rosas99/monster/internal/usercenter/model"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"regexp"
)

// Create implements the 'Create' method of the IBiz interface, responsible for creating a new user.
func (b *userBiz) Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {

	var userM model.UserM
	_ = copier.Copy(&userM, rq)
	if err := b.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return nil, errno.ErrUserAlreadyExist
		}
		return nil, err
	}

	return &v1.CreateUserResponse{UserID: userM.ID}, nil
}
