package message

import "github.com/rosas99/monster/internal/sms/service"

// MessageController represents the type that holds the controller state and behavior.
type MessageController struct {
	svc *service.SmsServerService
}

// New creates a new instance of the MessageController with the provided service layer.
// It returns a pointer to the newly created MessageController.
func New(svc *service.SmsServerService) *MessageController {
	return &MessageController{svc: svc}
}
