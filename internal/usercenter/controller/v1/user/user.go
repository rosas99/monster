package user

import "github.com/rosas99/monster/internal/usercenter/service"

type Controller struct {
	svc *service.UserCenterService
}

func New(svc *service.UserCenterService) *Controller {
	return &Controller{svc: svc}
}
