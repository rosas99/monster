package service

import (
	"context"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
)

// AILIYUNMessageReport is a method for receive message reports.
// It takes a AILIYUNReportListRequest as input and returns an CommonResponse or an error.
func (s *SmsServerService) AILIYUNMessageReport(ctx context.Context, rq *v1.AILIYUNReportListRequest) (*v1.CommonResponse, error) {
	log.C(ctx).Infow("AILIYUNMessageReport function called")
	return s.biz.Messages().AILIYUNReport(ctx, rq)
}

// SendMessage is a method for send a message.
// It takes a SendMessageRequest as input and returns an CommonResponse or an error.
func (s *SmsServerService) SendMessage(ctx context.Context, rq *v1.SendMessageRequest) (*v1.CommonResponse, error) {
	log.C(ctx).Infow("SendMessage function called")
	return s.biz.Messages().Send(ctx, rq)
}

// CodeVerify is a method for verify a verification code from message.
// It takes a VerifyCodeRequest as input and returns an CommonResponse or an error.
func (s *SmsServerService) CodeVerify(ctx context.Context, rq *v1.VerifyCodeRequest) (*v1.CommonResponse, error) {
	log.C(ctx).Infow("CreateOrder function called")
	return s.biz.Messages().CodeVerify(ctx, rq)
}
