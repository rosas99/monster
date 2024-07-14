package template

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

func (b *TemplateController) Get(c *gin.Context) {
	var r v1.GetTemplateRequest

	//param := c.Param("id")
	//r.Id = param
	//fmt.Print(r.Id)

	template, err := b.svc.GetTemplate(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
