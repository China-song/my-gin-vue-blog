package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	ginblog "my-gin-vue-blog/internal"
)

func main() {
	configPath := flag.String("c", "../config.yaml", "配置文件的路径")
	flag.Parse()

	// TODO: 初始化gorm对象
	_ = configPath

	r := gin.New()                      // 没有任何中间件
	r.Use(gin.Logger(), gin.Recovery()) // 全局中间件	 使用自带的日志和恢复中间件

	// 注册路由
	ginblog.RegisterHandlers(r)

	r.Run(":8765")
}
