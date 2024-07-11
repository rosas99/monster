package user

import (
	"context"
	"fmt"
	v1 "github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/token"
)

// ChangePassword 是 UserBiz 接口中 `ChangePassword` 方法的实现.
func (b *userBiz) Authorize(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginResponse, error) {
	// 只需要认证，在修改用户时才需要授权
	username, err := token.Parse(rq.Username, "config.key")
	fmt.Print(username)
	if err != nil {
		//core.WriteResponse(c, errno.ErrTokenInvalid, nil)
		//core.WriteResponse(ctx, nil, nil)

		return nil, err
	}
	return &v1.LoginResponse{}, nil
}
