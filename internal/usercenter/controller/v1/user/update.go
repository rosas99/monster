package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
)

func (b *Controller) Update(c *gin.Context) {
	_, _ = b.svc.Update(c, nil)

	core.WriteResponse(c, nil, "order")

}
