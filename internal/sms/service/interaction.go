package service

import (
	"context"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
)

func (s *SmsServerService) AILIYUNInteractionCallback(ctx context.Context, rq *v1.AILIYUNCallbackListRequest) (*v1.CommonResponse, error) {
	log.C(ctx).Infow("CreateOrder function called")
	return s.biz.Interaction().AILIYUNCallback(ctx, rq)
}
