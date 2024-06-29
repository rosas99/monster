package message

import "github.com/rosas99/monster/internal/sms/service"

type MessageController struct {
	svc *service.SmsServerService
}

func New(svc *service.SmsServerService) *MessageController {
	return &MessageController{svc: svc}
}
