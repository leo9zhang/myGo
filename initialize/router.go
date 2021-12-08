package initialize

import (
	"github.com/gin-gonic/gin"
	"myGo/global"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	global.GVA_LOG.Info("use middleware logger")

	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.GVA_LOG.Info("register swagger handler")

	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
