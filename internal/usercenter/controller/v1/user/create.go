// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	v1 "github.com/rosas99/monster/pkg/api/usercenter/v1"
)

func (b *Controller) Create(c *gin.Context) {
	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		// 使用错误码
		core.WriteResponse(c, err, nil)
		return
	}

	order, err := b.svc.Create(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, nil, order)

}
