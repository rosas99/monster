package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/known"
	"github.com/rosas99/monster/pkg/token"
)

// Authn 是认证中间件，用来从 gin.Context 中提取 token 并验证 token 是否合法，
// 如果合法则将 token 中的 sub 作为<用户名>存放在 gin.Context 的 XUsernameKey 键中.
func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
		username, err := token.ParseRequest(c)
		if err != nil {
			//core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			core.WriteResponse(c, nil, nil)
			c.Abort()

			return
		}

		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}
