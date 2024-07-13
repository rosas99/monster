package message

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// AILIYUNCallback is a controller for receive uplink messages from Alibaba Cloud.
func (b *MessageController) AILIYUNCallback(c *gin.Context) {
	var r v1.AILIYUNCallbackListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)

	}
	template, err := b.svc.AILIYUNInteractionCallback(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
