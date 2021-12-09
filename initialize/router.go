package initialize

import (
	"github.com/gin-gonic/gin"
	"myGo/global"
	"myGo/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	global.GVA_LOG.Info("use middleware logger")

	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.GVA_LOG.Info("register swagger handler")

	// 获取路由实例
	customerRouter := router.RouterGroupApp.Customer
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	PrivateGroup := Router.Group("")
	{
		customerRouter.InitCustomerRouter(PrivateGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
