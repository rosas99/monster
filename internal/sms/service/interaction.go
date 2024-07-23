package service

import (
	"context"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
)

// AILIYUNUplink is a method for receive an uplink message.
// It takes a AILIYUNUplinkListRequest as input and returns an error.
func (s *SmsServerService) AILIYUNUplink(ctx context.Context, rq *v1.AILIYUNUplinkListRequest) error {
	log.C(ctx).Infow("AILIYUNUplink function called")
	return s.biz.Interaction().AILIYUNUplink(ctx, rq)
}
