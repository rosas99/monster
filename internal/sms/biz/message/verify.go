package message

import (
	"context"
	"github.com/rosas99/monster/internal/pkg/errno"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// CodeVerify verifies a message code and clear cache if success.
func (b *messageBiz) CodeVerify(ctx context.Context, rq *v1.VerifyCodeRequest) error {

	key := factory.WrapperCode(rq.Mobile, rq.TemplateCode)
	code, err := b.rds.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if rq.Code != code {
		return errno.ErrBind
	}
	b.rds.Del(ctx, key)

	return nil

}
