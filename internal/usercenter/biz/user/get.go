package user

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"gorm.io/gorm"
)

// Get implements the 'Get' method of the IBiz interface, which retrieves user information based on the provided request.
func (b *userBiz) Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	filters := map[string]any{"user_name": rq.Username}
	userM, err := b.ds.Users().Fetch(ctx, filters)
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
