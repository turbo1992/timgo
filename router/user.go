package router

import (
	"github.com/gin-gonic/gin"
	v1 "tim-go/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("v1/user")
	{
		UserRouter.POST("register", v1.Register)
		UserRouter.POST("login", v1.Login)
		UserRouter.GET("detail", v1.UserDetail)
		UserRouter.POST("add", v1.AddUser)
		UserRouter.GET("find", v1.Find)
	}
}