package template

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
)

func (b *TemplateController) Update(c *gin.Context) {
	ret, err := b.svc.UpdateTemplate(c, nil)
	if err != nil {
		return
	}
	core.WriteResponse(c, nil, ret)

}
