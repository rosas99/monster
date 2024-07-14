package template

import "github.com/rosas99/monster/internal/sms/service"

// Controller represents the type that holds the controller state and behavior.
type Controller struct {
	svc *service.SmsServerService
}

// New creates a new instance of the Controller with the provided service layer.
// It returns a pointer to the newly created TemplateController.
func New(svc *service.SmsServerService) *Controller {
	return &Controller{svc: svc}
}
