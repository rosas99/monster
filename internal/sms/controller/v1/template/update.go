package template

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

func (b *Controller) Update(c *gin.Context) {
	var r v1.UpdateTemplateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	ret, err := b.svc.UpdateTemplate(c, c.Param("id"), &r)
	if err != nil {
		return
	}
	core.WriteResponse(c, nil, ret)

}
