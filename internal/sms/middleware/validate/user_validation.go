// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package validate

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/sms/store"
	"regexp"
	"strconv"
)

// Validation make sure users have the right resource permission and operation.
func Validation(ds store.IStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.FullPath() {
		case "/v1/template":
			if c.Request.Method == "GET" {
				id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
				_, err := ds.Templates().Get(context.Background(), id)
				if err != nil {
					// log kpi
					c.Abort()
					return
				}
			}
		default:
		}

		c.Next()
	}
}

func isMobileNo(mobiles string) bool {
	pattern := `^[0-9]{6}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(mobiles)
}
