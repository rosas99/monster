package interaction

import "github.com/rosas99/monster/internal/sms/service"

// InteractionController represents the type that holds the controller state and behavior.
type InteractionController struct {
	svc *service.SmsServerService
}

// New creates a new instance of the InteractionController with the provided service layer.
// It returns a pointer to the newly created InteractionController.
func New(svc *service.SmsServerService) *InteractionController {
	return &InteractionController{svc: svc}
}
