package message

import (
	"context"
	"github.com/rosas99/monster/internal/pkg/errno"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
)

// CodeVerify is used to verify the verification code entered by the user.
func (b *messageBiz) CodeVerify(ctx context.Context, rq *v1.VerifyCodeRequest) error {

	key := types.WrapperCode(rq.Mobile, rq.TemplateCode)
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
