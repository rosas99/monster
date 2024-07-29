package user

import (
	"context"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

// Update implements the 'Update' method of the IBiz interface, allowing modification of user data based on the provided update request.
func (b *userBiz) Update(ctx context.Context, rq *v1.UpdateUserRequest) error {
	filters := map[string]any{"user_name": rq.Username}
	userM, err := b.ds.Users().Fetch(ctx, filters)
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
