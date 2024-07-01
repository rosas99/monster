package service

import (
	"context"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SmsServerService) CreateTemplate(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {
	log.C(ctx).Infow("CreateOrder function called")
	return s.biz.Templates().Create(ctx, rq)
}

func (s *SmsServerService) ListTemplate(ctx context.Context, rq *v1.ListTemplateRequest) (*v1.ListTemplateResponse, error) {
	return s.biz.Templates().List(ctx, rq)
}

func (s *SmsServerService) GetTemplate(ctx context.Context, rq *v1.GetTemplateRequest) (*v1.TemplateReply, error) {
	log.C(ctx).Infow("GetOrder function called")
	return s.biz.Templates().Get(ctx, rq)
}

func (s *SmsServerService) UpdateTemplate(ctx context.Context, rq *v1.UpdateTemplateRequest) (*emptypb.Empty, error) {
	log.C(ctx).Infow("UpdateOrder function called")
	return &emptypb.Empty{}, s.biz.Templates().Update(ctx, rq)
}

func (s *SmsServerService) DeleteTemplate(ctx context.Context, rq *v1.DeleteTemplateRequest) (*emptypb.Empty, error) {
	log.C(ctx).Infow("DeleteOrder function called")
	return &emptypb.Empty{}, s.biz.Templates().Delete(ctx, rq)
}
