package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	v1 "github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/log"
)

func (b *Controller) ChangePassword(c *gin.Context) {
	log.C(c).Infow("Change password function called")

	// todo param放到struct 中
	var r v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		//core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	//if _, err := govalidator.ValidateStruct(r); err != nil {
	//	//core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
	//
	//	return
	//}

	if err := b.svc.ChangePassword(c, &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)

}
