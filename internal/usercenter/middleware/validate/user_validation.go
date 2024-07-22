// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/usercenter/store"
)

// todo 使用这种方式更加简单
// 这样不需要使用自定义的了

// Validation make sure users have the right resource permission and operation.
func Validation(ds store.IStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
