package router

import (
	"GolangStudy/Practice/shop/user-web/api"
	"GolangStudy/Practice/shop/user-web/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	zap.S().Info("配置用户相关的url")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
	}
}
