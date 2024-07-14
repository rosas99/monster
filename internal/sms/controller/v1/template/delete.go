package template

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
)

func (b *Controller) Delete(c *gin.Context) {

	template, err := b.svc.DeleteTemplate(c, c.Param("id"))
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
