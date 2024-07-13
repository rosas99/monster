package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

func (b *Controller) Delete(c *gin.Context) {
	var r v1.DeleteUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
	}
	template, err := b.svc.Delete(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}

	//if _, err := ctrl.a.RemoveNamedPolicy("p", username, "", ""); err != nil {
	//	core.WriteResponse(c, err, nil)
	//
	//	return
	//}
	core.WriteResponse(c, nil, template)

}
