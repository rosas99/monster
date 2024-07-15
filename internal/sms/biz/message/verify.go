package message

import (
	"context"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// CodeVerify verifies a message code and clear cache if success.
func (b *messageBiz) CodeVerify(ctx context.Context, rq *v1.VerifyCodeRequest) error {

	key := factory.WrapperCode(rq.Mobile, rq.TemplateCode)
	code, err := b.rds.Get(ctx, key).Result()
	if rq.Code != code {
		return err
	}
	b.rds.Del(ctx, key)
	message := map[string]any{
		"test":  "value1",
		"other": 123,
	}
	b.logger.LogKpi(message)
	return nil

}
