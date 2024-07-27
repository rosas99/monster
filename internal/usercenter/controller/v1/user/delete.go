package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

// Delete handles the deletion of a user within the context of the Controller.
func (b *Controller) Delete(c *gin.Context) {
	var r v1.DeleteUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
	}
	_, err := b.svc.Delete(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}

	_, err = b.svc.Auth.RemoveNamedPolicy("p", r.Username, "", "")
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)

}
