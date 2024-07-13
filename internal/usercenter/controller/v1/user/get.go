package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

func (b *Controller) Get(c *gin.Context) {
	var r v1.GetUserRequest

	//param := c.Param("id")
	//r.Id = param
	//fmt.Print(r.Id)

	template, err := b.svc.Get(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
