package main

import (
	"github.com/gin-gonic/gin"
	"github.com/learn/gostudy/config"
	"github.com/learn/gostudy/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}
