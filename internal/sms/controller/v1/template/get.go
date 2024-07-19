package template

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"strconv"
)

func (b *Controller) Get(c *gin.Context) {
	i, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	template, err := b.svc.GetTemplate(c, i)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
