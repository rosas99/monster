package message

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

func (b *Controller) CodeVerify(c *gin.Context) {
	var r v1.VerifyCodeRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
	}

	err := b.svc.CodeVerify(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}

	core.WriteResponse(c, nil, nil)

}
