package user

import (
	"context"
	"errors"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/auth"
	"github.com/rosas99/monster/pkg/log"
)

// ChangePassword implements the 'ChangePassword' method of the IBiz interface.
func (b *userBiz) ChangePassword(ctx context.Context, username string, rq *v1.ChangePasswordRequest) error {
	filters := map[string]any{"user_name": username}
	userM, err := b.ds.Users().Fetch(ctx, filters)
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
