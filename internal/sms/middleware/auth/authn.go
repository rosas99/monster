package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/errno"
	"github.com/rosas99/monster/internal/pkg/known"
	"github.com/rosas99/monster/internal/pkg/middleware/auth"
	jwtutil "github.com/rosas99/monster/internal/pkg/util/jwt"
)

// BasicAuth creates a middleware that authenticates requests using the provided AuthProvider.
func BasicAuth(a auth.AuthProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := jwtutil.TokenFromServerContext(c)
		//todo 补充注释

		userID, err := a.Auth(c.Request.Context(), accessToken)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set(known.UsernameKey, userID)
		c.Next()
	}
}
