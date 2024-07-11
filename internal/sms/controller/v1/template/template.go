package template

import "github.com/rosas99/monster/internal/sms/service"

// TemplateController represents the type that holds the controller state and behavior.
type TemplateController struct {
	svc *service.SmsServerService
}

// New creates a new instance of the TemplateController with the provided service layer.
// It returns a pointer to the newly created TemplateController.
func New(svc *service.SmsServerService) *TemplateController {
	return &TemplateController{svc: svc}
}
