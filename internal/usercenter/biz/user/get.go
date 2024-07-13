package user

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"gorm.io/gorm"
)

func (b *userBiz) Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	userM, err := b.ds.Users().Get(ctx, rq.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	var resp v1.GetUserResponse
	_ = copier.Copy(&resp, userM)
	resp.CreatedAt = userM.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = userM.UpdatedAt.Format("2006-01-02 15:04:05")

	return &resp, nil
}
