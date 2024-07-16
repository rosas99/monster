package user

import (
	"context"
	"fmt"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/token"
)

// ChangePassword 是 IBiz 接口中 `ChangePassword` 方法的实现.
func (b *userBiz) Authorize(ctx context.Context, rq *v1.LoginRequest2) (*v1.LoginResponse2, error) {
	username, err := token.Parse(rq.Username, "config.key")
	fmt.Print(username)
	if err != nil {
		return nil, err
	}
	return &v1.LoginResponse2{}, nil
}
