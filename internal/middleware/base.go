package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

// WithRedisDB 将 redis.Client 注入到 gin.Context
// handler 中通过 c.MustGet(g.CTX_RDB).(*redis.Client) 来使用
func WithRedisDB(rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(global.CTX_RDB, rdb)
		ctx.Next()
	}
}

// WithCookieStore 基于cookie的session
func WithCookieStore(name, secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{Path: "/", MaxAge: 600})
	return sessions.Sessions(name, store)
}
