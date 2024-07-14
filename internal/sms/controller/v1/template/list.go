package template

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

func (b *Controller) List(c *gin.Context) {

	var r v1.ListTemplateRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, err, nil)
	}

	template, err := b.svc.ListTemplate(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
