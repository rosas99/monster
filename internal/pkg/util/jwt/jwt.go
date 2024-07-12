package jwt

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	// bearerWord the bearer key word for authorization.
	bearerWord string = "Bearer"

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	authorizationKey string = "Authorization"
)

func TokenFromServerContext(c *gin.Context) string {

	auths := strings.SplitN(c.Request.Header.Get(authorizationKey), " ", 2)
	if len(auths) == 2 && strings.EqualFold(auths[0], bearerWord) {
		return auths[1]
	}
	return ""
}
