package usercenter

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
	"github.com/rosas99/monster/internal/pkg/errno"
	"github.com/rosas99/monster/internal/usercenter/controller/v1/user"
	mwauth "github.com/rosas99/monster/internal/usercenter/middleware/auth"
	"github.com/rosas99/monster/internal/usercenter/service"
	"github.com/rosas99/monster/internal/usercenter/store/mysql"
	"github.com/rosas99/monster/pkg/auth"
)

func installRouters(g *gin.Engine, svc *service.UserCenterService) {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 pprof 路由
	pprof.Register(g)

	authz, err := auth.NewAuthz(mysql.S.DB())
	if err != nil {
		return
	}

	uc := user.New(svc)
	g.POST("/login", uc.Login)

	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/user")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mwauth.Authn(), mwauth.Authz(authz))
			userv1.GET(":name", uc.Get)
			userv1.PUT(":name", uc.Update)
			userv1.GET("", uc.List)
			userv1.DELETE(":name", uc.Delete)
		}
	}

}
