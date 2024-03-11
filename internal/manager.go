package ginblog

import (
	"github.com/gin-gonic/gin"
	"my-gin-vue-blog/internal/handle"
)

var (
	userAuthAPI handle.UserAuth // 用户账号
)

func RegisterHandlers(r *gin.Engine) {
	registerBaseHandler(r)
}

func registerBaseHandler(r *gin.Engine) {
	base := r.Group("/api")

	base.POST("/login", userAuthAPI.Login)
}
