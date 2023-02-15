package router

import (
	v1 "gin-blog/api/v1"
	"gin-blog/middleware"
	"gin-blog/utils"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() {
	// 设置项目模式
	gin.SetMode(utils.AppMode)
	// 初始化路由
	router := gin.New()
	router.Use(gin.Recovery())
	// 配置跨域中间件
	router.Use(middleware.Cors())
	// 使用日志中间件
	router.Use(middleware.Logger())
	// 获取userapi
	userApi := v1.UserApi{}
	// 获取catapi
	var cateApi v1.CategoryApi
	// 获取articleapi
	var artcileApi v1.ArticleApi
	// loginapi
	loginApi := v1.LoginApi{}
	// uploadapi
	var uploadApi v1.UploadApi
	// 无需鉴权
	NoRouter := router.Group("/v1")

	{
		// 后台登录
		NoRouter.POST("admin/login", loginApi.UserLogin)
		// 添加用户
		NoRouter.POST("user/add", userApi.Add)
		// 查询用户列表
		NoRouter.GET("user/list", userApi.List)
		// 查询分类列表
		NoRouter.GET("category/list", cateApi.ListCate)
		// 查询文章
		NoRouter.GET("article/only/:id", artcileApi.OnlyArt)
		NoRouter.GET("article/catelist/:cid", artcileApi.CateList)
		NoRouter.GET("article/list", artcileApi.ListArt)
	}
	// 需要鉴权的路由
	JwtRouter := router.Group("/v1")
	// 使用中间件
	JwtRouter.Use(middleware.JwtToken())
	// 使用cookie中间件
	JwtRouter.Use(middleware.Cookier())
	{

		// 删除用户
		JwtRouter.DELETE("user/delete/:id", userApi.Delete)
		// 编辑用户
		JwtRouter.PUT("user/edit/:id", userApi.Edit)

		// 增加分类
		JwtRouter.POST("category/add", cateApi.AddCate)
		// 删除分类
		JwtRouter.DELETE("category/delete/:id", cateApi.DeleteCate)
		// 修改分类
		JwtRouter.PUT("category/edit/:id", cateApi.EditCate)

		// 增加文章
		JwtRouter.POST("article/add", artcileApi.AddArt)
		// 删除文章
		JwtRouter.DELETE("article/delete/:id", artcileApi.DeleteArt)
		// 修改文章
		JwtRouter.PUT("article/edit/:id", artcileApi.EditArt)

		// 上传图片到七牛云
		JwtRouter.POST("/upload/img", uploadApi.UploadImg)

	}
	router.Run(utils.HttpPort)
}
