package message

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/errno"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
)

// AILIYUNCallback is a controller for receive uplink messages from Alibaba Cloud.
func (b *Controller) AILIYUNCallback(c *gin.Context) {
	defer core.WriteResponse(c, errno.AiliCloudSuccess, nil)

	var r v1.AILIYUNUplinkListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.C(c).Errorf("Error occurred while binding the request body to the struct. %s", err)
		return
	}

	err := validator.New().Struct(r)
	if err != nil {
		log.C(c).Errorf("Parameter verification failed. %s", err)
		return
	}

	_ = b.svc.AILIYUNUplink(c, &r)

}
