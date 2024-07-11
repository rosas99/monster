package sms

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/client/usercenter"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/sms/controller/v1/interaction"
	"github.com/rosas99/monster/internal/sms/controller/v1/message"
	"github.com/rosas99/monster/internal/sms/controller/v1/template"
	mw "github.com/rosas99/monster/internal/sms/middleware/auth"
	"github.com/rosas99/monster/internal/sms/service"
	v1api "github.com/rosas99/monster/pkg/api/sms/v1"
)

func installRouters(g *gin.Engine, svc *service.SmsServerService) {
	// register 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, v1api.ErrorOrderAlreadyExists("route not found"), nil)
	})

	// register pprof handler
	pprof.Register(g)

	// creates a user center client by GRPC
	impl := usercenter.NewUserCenterServer()
	// creates a v1 router group and adds an auth middleware.
	v1 := g.Group("/v1", mw.BasicAuth(impl))
	{
		// create template router group
		templatev1 := v1.Group("/template")
		{
			tl := template.New(svc)
			templatev1.POST("/create", tl.Create)
			templatev1.POST("/update", tl.Update)
			templatev1.GET("/:id", tl.Get)
			templatev1.POST("/getList", tl.List)
			templatev1.POST("/delete", tl.Delete)
		}

		// creates message router group
		msgv1 := v1.Group("/message")
		{
			ms := message.New(svc)
			msgv1.POST("/send", ms.Send)
			msgv1.POST("/verify", ms.CodeVerify)

			// todo 需要支持公网ip
			msgv1.POST("/report/ailiyun", ms.AiliReport)

		}

		// creates interaction router group
		itv1 := v1.Group("/interaction")
		{
			it := interaction.New(svc)
			// todo 需要支持公网ip
			itv1.POST("/ailiyun", it.AILIYUNCallback)
		}

	}

}
