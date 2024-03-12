package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my-gin-vue-blog/internal/global"
)

// WithGormDB 将 gorm.DB 注入到 gin.Context
// handler 中通过 c.MustGet(g.CTX_DB).(*gorm.DB) 来使用
func WithGormDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(global.CTX_DB, db)
		ctx.Next()
	}
}
