package message

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/known"
	"github.com/rosas99/monster/internal/sms/monitor"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"time"

	// custom gin validators.
	_ "github.com/rosas99/monster/pkg/validator"
)

func (b *Controller) CodeVerify(c *gin.Context) {
	start := time.Now().UnixMilli()

	var r v1.VerifyCodeRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
	}

	err := b.svc.CodeVerify(c, &r)
	if err != nil {
		monitor.GetMonitor().LogKpi(
			"VerifyCode Message",
			c.Request.Header.Get(known.TraceIDKey),
			r.TemplateCode,
			false,
			time.Now().UnixMilli()-start,
		)
		core.WriteResponse(c, err, nil)
	}

	core.WriteResponse(c, nil, nil)

}
