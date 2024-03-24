package ginblog

import (
	"github.com/gin-gonic/gin"
	"my-gin-vue-blog/docs"
	"my-gin-vue-blog/internal/handle"
	"my-gin-vue-blog/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	userAuthAPI handle.UserAuth // 用户账号
)

func RegisterHandlers(r *gin.Engine) {
	// swagger
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerBaseHandler(r)
}

func registerBaseHandler(r *gin.Engine) {
	base := r.Group("/api")

	base.POST("/login", userAuthAPI.Login)
}

func registerAdminHandler(r *gin.Engine) {
	auth := r.Group("/api")

	// 加中间件，如 JWT验证？
	auth.Use(middleware.JWTAuth())
	auth.Use(middleware.PermissionCheck())
}
