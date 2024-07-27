package user

import (
	"context"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

// Delete implements the 'Delete' method of the OrderBiz interface.
func (b *userBiz) Delete(ctx context.Context, rq *v1.DeleteUserRequest) error {
	if err := b.ds.Users().Delete(ctx, rq.Username); err != nil {
		return err
	}

	return nil
}
