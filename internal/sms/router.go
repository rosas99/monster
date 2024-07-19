package sms

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/errno"
	"github.com/rosas99/monster/internal/sms/controller/v1/message"
	"github.com/rosas99/monster/internal/sms/controller/v1/template"
	"github.com/rosas99/monster/internal/sms/service"
)

func installRouters(g *gin.Engine, svc *service.SmsServerService) {
	// register 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// register pprof handler
	pprof.Register(g)

	// creates a v1 router group and adds an auth middleware.
	// get a grpc usercenter client
	// todo 开启认证
	// v1 := g.Group("/v1", mw.BasicAuth(usercenter.GetClient()))
	v1 := g.Group("/v1")
	{
		v1.Use()
		// create template router group
		templatev1 := v1.Group("/template")
		{
			tl := template.New(svc)
			templatev1.POST("", tl.Create)
			templatev1.PUT("", tl.Update)
			templatev1.GET("/:id", tl.Get)
			templatev1.GET("", tl.List)
			templatev1.DELETE("/:id", tl.Delete)
		}

		// creates message router group
		msgv1 := v1.Group("/message")
		{
			ms := message.New(svc)
			msgv1.POST("/send", ms.Send)
			msgv1.POST("/verify", ms.CodeVerify)

			// todo 需要支持公网ip
			msgv1.POST("/report/ailiyun", ms.AiliReport)
			msgv1.POST("/interaction/ailiyun", ms.AILIYUNCallback)

		}

	}

}
