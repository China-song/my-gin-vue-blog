package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	ginblog "my-gin-vue-blog/internal"
	"my-gin-vue-blog/internal/global"
	"my-gin-vue-blog/internal/middleware"
)

// @title           My Gin-Vue-Blog API
// @version         1.0
// @description     博客后端

// @host      localhost:8765
func main() {
	configPath := flag.String("c", "../config.yaml", "配置文件的路径")
	flag.Parse()

	conf := global.ReadConfig(*configPath)
	// 得到gorm.DB对象
	db := ginblog.InitDatabase(conf)
	rdb := ginblog.InitRedis(conf)

	r := gin.New()                      // 没有任何中间件
	r.Use(gin.Logger(), gin.Recovery()) // 全局中间件	 使用自带的日志和恢复中间件

	r.Use(middleware.WithGormDB(db))
	r.Use(middleware.WithRedisDB(rdb))
	r.Use(middleware.WithCookieStore(conf.Session.Name, conf.Session.Salt))

	// 注册路由
	ginblog.RegisterHandlers(r)

	serverAddr := conf.Server.Port

	r.Run(serverAddr)
}
