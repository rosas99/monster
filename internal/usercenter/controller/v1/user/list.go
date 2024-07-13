package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

func (b *Controller) List(c *gin.Context) {
	// query 示例
	//if err := c.ShouldBindQuery(&r); err != nil {
	//	core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
	//
	//	return
	//}

	var r v1.ListUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
	}

	template, err := b.svc.List(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
