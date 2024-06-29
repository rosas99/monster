// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package validate

import (
	"context"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/sms/store"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"regexp"

	"github.com/gin-gonic/gin"
)

// todo 使用这种方式更加简单
// 这样不需要使用自定义的了

// Validation make sure users have the right resource permission and operation.
func Validation(ds store.IStore) gin.HandlerFunc {
	return func(c *gin.Context) {

		switch c.FullPath() {
		// todo 根据url校验：
		// 参数非空 模板校验 手机号规则校验
		case "/v1/users":
			_, err := ds.Templates().Get(context.Background(), "")
			if err != nil {
				// log kpi
				// 返回错误码
				c.Abort()
				return
			}
			if !isMobileNo(c.GetString("mobile")) {
				// log kpi
				// 返回错误码
				c.Abort()
				return
			}

		// 手机号白名单校验 修改为在mq消费前校验
		case "/v1/users/:name", "/v1/users/:name/change_password":
			var r v1.CreateTemplateRequest
			if err := c.ShouldBindJSON(&r); err != nil {
				core.WriteResponse(c, err, nil)
				// todo 了解gin如何返回错误 如：
				/*
					c.JSON(http.StatusBadRequest, gin.H{
							"err": err.Error(),
						})
				*/
			}

		default:
		}

		c.Next()
	}
}

func isMobileNo(mobiles string) bool {
	// 定义一个正则表达式，匹配6位数字
	pattern := `^[0-9]{6}$`
	// 编译正则表达式
	re := regexp.MustCompile(pattern)
	// 检查手机号码是否符合正则表达式
	return re.MatchString(mobiles)
}
