package user

import (
	"context"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

// Delete 是 OrderBiz 接口中 `Delete` 方法的实现.
func (b *userBiz) Delete(ctx context.Context, rq *v1.DeleteUserRequest) error {
	if err := b.ds.Users().Delete(ctx, rq.UserId); err != nil {
		return err
	}

	return nil
}
