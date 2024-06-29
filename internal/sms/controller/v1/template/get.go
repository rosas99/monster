// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package template

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

type validator interface {
	Validate() error
}

func (b *TemplateController) Get(c *gin.Context) {
	var r v1.GetTemplateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
	}

	if err := r.Validate(); err != nil {
		core.WriteResponse(c, err, nil)
	}

	template, err := b.svc.GetTemplate(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

	}
	core.WriteResponse(c, nil, template)

}
