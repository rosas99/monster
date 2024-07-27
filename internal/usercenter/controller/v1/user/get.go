package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/pkg/api/usercenter/v1"
)

// Get retrieves user information within the context of the Controller.
func (b *Controller) Get(c *gin.Context) {
	var r v1.GetUserRequest
	template, err := b.svc.Get(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}

	core.WriteResponse(c, nil, template)

}
