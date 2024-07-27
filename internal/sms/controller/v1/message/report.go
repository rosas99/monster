package message

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/errno"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
)

// AiliReport handles the request for aili cloud message reports.
func (b *Controller) AiliReport(c *gin.Context) {
	defer core.WriteResponse(c, errno.AiliCloudSuccess, nil)

	var r v1.AILIYUNReportListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.C(c).Errorf("Error occurred while binding the request body to the struct: %v", err)
	}
	err := b.svc.AILIYUNMessageReport(c, &r)
	if err != nil {
		log.C(c).Errorf("Error occurred while processing AILIYUNMessageReport: %v", err)
	}

}
