package template

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
)

func (b *TemplateController) Update(c *gin.Context) {
	_, _ = b.svc.GetTemplate(c, nil)

	// todo 临时测试用
	core.WriteResponse(c, nil, "order")

}
