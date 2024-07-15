package message

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/errno"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// AILIYUNCallback is a controller for receive uplink messages from Alibaba Cloud.
func (b *Controller) AILIYUNCallback(c *gin.Context) {
	var r v1.AILIYUNUplinkListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)

	}
	err := b.svc.AILIYUNUplink(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, errno.AiliCloudSuccess, nil)

}
