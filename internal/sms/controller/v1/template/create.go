package template

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	// custom gin validators.
	_ "github.com/rosas99/monster/pkg/validator"
)

func (b *Controller) Create(c *gin.Context) {
	var r v1.CreateTemplateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	order, err := b.svc.CreateTemplate(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, nil, order)

}
