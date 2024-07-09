// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package sms

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	mw "github.com/rosas99/monster/internal/pkg/middleware/auth"
	"github.com/rosas99/monster/internal/usercenter/controller/v1/user"
	"github.com/rosas99/monster/internal/usercenter/service"
	"github.com/rosas99/monster/internal/usercenter/store/mysql"
	v1api "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/auth"
)

func installRouters(g *gin.Engine, svc *service.UserCenterService) {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, v1api.ErrorOrderAlreadyExists("route not found"), nil)
	})

	// 注册 pprof 路由
	pprof.Register(g)

	// 创建 v1 路由分组，并添加认证中间件
	//v1 := g.Group("/v1", mw.BasicAuth(accounts))

	authz, err := auth.NewAuthz(mysql.S.DB())
	if err != nil {
		return
	}

	v1 := g.Group("/v1")
	{
		uc := user.New(svc)
		userv1 := v1.Group("/msg")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mw.Authn(), mw.Authz(authz))
			userv1.POST("", uc.Create)
			userv1.POST("", uc.Create)
			userv1.POST("", uc.Create)
		}

	}

}
