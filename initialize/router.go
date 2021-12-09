package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"myGo/global"
	"myGo/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//ginSwagger.WrapHandler(swaggerFiles.Handler,
	//	ginSwagger.URL("http://localhost:8888/swagger/doc.json"),
	//	ginSwagger.DefaultModelsExpandDepth(-1))
	//global.Logger.Info("register swagger handle")

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
	global.Logger.Info("router register success")
	return Router
}
