package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

// List retrieves a list of users within the context of the Controller.
func (b *Controller) List(c *gin.Context) {
	var r v1.ListUserRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, err, nil)
	}

	template, err := b.svc.List(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
