package router

import (
	v1 "gin-blog/api/v1"
	"gin-blog/utils"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() {
	// 设置项目模式
	gin.SetMode(utils.AppMode)
	// 初始化路由
	router := gin.Default()
	// 路由分组
	UserRouter := router.Group("user")
	// 获取userapi
	userApi := &v1.UserApi{}
	{
		// 查询用户列表
		UserRouter.GET("list", userApi.GetList)
		// 添加用户
		UserRouter.POST("add", userApi.Add)
	}
	router.Run(utils.HttpPort)
}
