package message

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/errno"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// AiliReport handles the request for aili cloud message reports.
func (b *Controller) AiliReport(c *gin.Context) {
	var r v1.AILIYUNReportListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)

	}
	err := b.svc.AILIYUNMessageReport(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	// 返回success，否则会自动重试
	core.WriteResponse(c, errno.AiliCloudSuccess, nil)

}
