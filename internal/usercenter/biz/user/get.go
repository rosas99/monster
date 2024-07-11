package user

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	v1 "github.com/rosas99/monster/pkg/api/usercenter/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (b *userBiz) Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	userM, err := b.ds.Users().Get(ctx, rq.GetUserName())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	var resp v1.GetUserResponse
	_ = copier.Copy(&resp, userM)
	resp.UserInfo.CreatedAt = timestamppb.New(userM.CreatedAt)
	resp.UserInfo.UpdatedAt = timestamppb.New(userM.UpdatedAt)

	return &resp, nil
}
