package service

import (
	"context"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
)

// AILIYUNInteractionCallback is a method for receive an uplink message.
// It takes a AILIYUNCallbackListRequest as input and returns an CommonResponse or an error.
func (s *SmsServerService) AILIYUNInteractionCallback(ctx context.Context, rq *v1.AILIYUNCallbackListRequest) (*v1.CommonResponse, error) {
	log.C(ctx).Infow("CreateOrder function called")
	return s.biz.Interaction().AILIYUNCallback(ctx, rq)
}
