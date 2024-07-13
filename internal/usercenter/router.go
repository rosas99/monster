package usercenter

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/usercenter/controller/v1/user"
	"github.com/rosas99/monster/internal/usercenter/service"
	v1api "github.com/rosas99/monster/pkg/api/sms/v1"
)

func installRouters(g *gin.Engine, svc *service.UserCenterService) {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, v1api.ErrorOrderAlreadyExists("route not found"), nil)
	})

	// 注册 pprof 路由
	pprof.Register(g)

	//authz, err := auth.NewAuthz(mysql.S.DB())
	//if err != nil {
	//	return
	//}

	v1 := g.Group("/v1")
	{
		uc := user.New(svc)
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			//userv1.Use(mw.Authn(), mw.Authz(authz))
			userv1.GET(":name1", uc.Get)
			userv1.PUT(":name", uc.Update)
			userv1.GET("", uc.List)
			userv1.DELETE(":name", uc.Delete)
		}
	}

}
