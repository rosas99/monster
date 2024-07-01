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

func (b *Controller) Get(c *gin.Context) {
	var r v1.GetUserRequest

	//param := c.Param("id")
	//r.Id = param
	//fmt.Print(r.Id)

	// todo 可修改为go playground
	if err := r.Validate(); err != nil {
		core.WriteResponse(c, err, nil)
	}

	template, err := b.svc.Get(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
