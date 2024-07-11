package user

import (
	"context"
	"errors"
	v1 "github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/auth"
	"github.com/rosas99/monster/pkg/log"
)

// ChangePassword 是 UserBiz 接口中 `ChangePassword` 方法的实现.
func (b *userBiz) ChangePassword(ctx context.Context, rq *v1.ChangePasswordRequest) error {
	userM, err := b.ds.Users().Get(ctx, rq.Username)
	if err != nil {
		return err
	}

	if err := auth.Compare(userM.Password, rq.OldPassword); err != nil {
		log.C(ctx).Errorw(err, "Failed to list orders from storage")

		return errors.New("old password is invalid")
	}

	userM.Password, _ = auth.Encrypt(rq.NewPassword)
	if err := b.ds.Users().Update(ctx, userM); err != nil {
		return err
	}

	return nil
}
