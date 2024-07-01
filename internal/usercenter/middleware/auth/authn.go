// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/known"
	"github.com/rosas99/monster/internal/pkg/middleware/auth"
	jwtutil "github.com/rosas99/monster/internal/pkg/util/jwt"
)

// middleware demo
// todo jwt认证
func BasicAuth(a auth.AuthProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := jwtutil.TokenFromServerContext(c)
		// rpc调用usercenter
		userID, err := a.Auth(c.Request.Context(), accessToken)
		if err != nil {

			// 返回错误码
			core.WriteResponse(c, err, nil)
			//core.WriteResponse(c, errno.ErrTokenInvalid, nil) // 具体错误码类型

			// 中断处理
			c.Abort()
			return
		}
		// todo 设置请求头响应参数等
		c.Set(known.UsernameKey, userID)
		c.Next()
	}
}
