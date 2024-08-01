package message

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/errno"
	"github.com/rosas99/monster/internal/pkg/known"
	"github.com/rosas99/monster/internal/sms/monitor"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"time"
)

func (b *Controller) Send(c *gin.Context) {
	start := time.Now().UnixMilli()
	var r v1.SendMessageRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	err := b.svc.SendMessage(c, &r)
	if err != nil {
		httpStatus, code, message := errno.Decode(err)

		monitor.GetMonitor().LogKpi(
			"Template Message",
			c.Request.Header.Get(known.TraceIDKey),
			r.TemplateCode,
			false,
			time.Now().UnixMilli()-start,
		)
		core.WriteResponse(c, &errno.Errno{HTTP: httpStatus, Code: code, Message: message}, nil)
		return
	}
	core.WriteResponse(c, errno.AiliCloudSuccess, nil)

}
