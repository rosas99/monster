package user

import (
	"context"
	v1 "github.com/rosas99/monster/pkg/api/usercenter/v1"
)

func (b *userBiz) Update(ctx context.Context, rq *v1.UpdateUserRequest) error {
	userM, err := b.ds.Users().Get(ctx, rq.GetUsername())
	if err != nil {
		return err
	}

	if rq.Email != nil {
		userM.Email = *rq.Email
	}

	if rq.Nickname != nil {
		userM.Nickname = *rq.Nickname
	}

	if rq.Phone != nil {
		userM.Phone = *rq.Phone
	}

	err = b.ds.Users().Update(ctx, userM)

	return err
}
