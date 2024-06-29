package template

import "github.com/rosas99/monster/internal/sms/service"

type TemplateController struct {
	svc *service.SmsServerService
}

func New(svc *service.SmsServerService) *TemplateController {
	return &TemplateController{svc: svc}
}
