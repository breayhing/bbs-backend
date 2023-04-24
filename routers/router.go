package routers

import (
	"github.com/gin-gonic/gin"
	v1 "zhongzhu-bbs/api/v1"
	"zhongzhu-bbs/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//User模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//Category模块路由接口

		//Article模块路由接口
	}

	r.Run(utils.HttpPort)
}
