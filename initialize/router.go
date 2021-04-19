package initialize

import (
	"github.com/gin-gonic/gin"
	"tim-go/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(gin.Recovery())
	// 打开就能玩https了
	// Router.Use(middleware.LoadTls())

	// 跨域
	//Router.Use(middleware.Cors())

	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("api")
	router.InitUserRouter(ApiGroup)                  // 注册用户路由

	return Router
}