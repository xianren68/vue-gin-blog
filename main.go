package main

import (
	"gin-blog/model"
	"gin-blog/router"
)

func main() {
	// 模型初始化
	model.InitDb()
	// 路由初始化
	router.InitRouter()
}
