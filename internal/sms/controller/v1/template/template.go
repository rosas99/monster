package template

import "github.com/rosas99/monster/internal/sms/service"

type Controller struct {
	svc *service.SmsServerService
}

func New(svc *service.SmsServerService) *Controller {
	return &Controller{svc: svc}
}
