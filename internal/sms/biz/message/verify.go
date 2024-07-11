package message

import (
	"context"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

func (b *messageBiz) CodeVerify(ctx context.Context, rq *v1.VerifyCodeRequest) (*v1.CommonResponse, error) {

	key := factory.WrapperCode(rq.Mobile, rq.TemplateCode)
	code, err := b.rds.Get(ctx, key).Result()
	if rq.Code != code {
		return &v1.CommonResponse{Code: 500, Msg: "fail"}, err
	}
	b.rds.Del(ctx, key)
	message := map[string]any{
		"test":  "value1",
		"other": 123,
	}
	b.logger.LogKpi(message)
	return &v1.CommonResponse{Code: 0, Msg: "success"}, nil

}
