package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/errno"
	v1 "github.com/rosas99/monster/pkg/api/usercenter/v1"
)

// Update handles the updating of a user's information within the context of the Controller.
func (b *Controller) Update(c *gin.Context) {
	var r v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	_, _ = b.svc.Update(c, &r)
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, "order")

}
