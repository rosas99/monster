package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
)

func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		c.JSON(errors.Code(err), err)
		return
	}

	c.JSON(http.StatusOK, data)
}
