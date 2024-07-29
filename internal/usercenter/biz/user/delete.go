package user

import (
	"context"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

// Delete implements the 'Delete' method of the OrderBiz interface.
func (b *userBiz) Delete(ctx context.Context, rq *v1.DeleteUserRequest) error {
	filters := map[string]any{"user_name": rq.Username}
	if err := b.ds.Users().Delete(ctx, filters); err != nil {
		return err
	}

	return nil
}
