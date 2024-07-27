package message

import (
	"context"
	"github.com/rosas99/monster/internal/pkg/errno"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
)

// todo 补充注释

// CodeVerify verifies a message code and clear cache if success.
func (b *messageBiz) CodeVerify(ctx context.Context, rq *v1.VerifyCodeRequest) error {

	key := factory.WrapperCode(rq.Mobile, rq.TemplateCode)
	log.C(ctx).Infof("Retrieving verification code for mobile: %s with template code: %s", rq.Mobile, rq.TemplateCode)

	code, err := b.rds.Get(ctx, key).Result()
	if err != nil {
		log.C(ctx).Errorf("Failed to retrieve code from cache with key: %s. Error: %v", key, err)
		return err
	}

	if rq.Code != code {
		log.C(ctx).Warnf("Verification failed for mobile: %s. Provided code does not match cached code.", rq.Mobile)
		return errno.ErrBind
	}
	log.C(ctx).Infof("Verification successful for mobile: %s", rq.Mobile)

	b.rds.Del(ctx, key)
	log.C(ctx).Infof("Deleted verification code for mobile: %s", rq.Mobile)

	return nil

}
