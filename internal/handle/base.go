package handle

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log/slog"
	"my-gin-vue-blog/internal/global"
	"my-gin-vue-blog/internal/model"
	"net/http"
)

// Response 响应结构体
type Response[T any] struct {
	Code    int    `json:"code"`    // 业务状态码
	Message string `json:"message"` // 响应消息
	Data    T      `json:"data"`    // 响应数据
}

// HTTP 码 + 业务码 + 消息 + 数据
func ReturnHttpResponse(c *gin.Context, httpCode, code int, msg string, data any) {
	c.JSON(httpCode, Response[any]{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

// 业务码 + 数据
func ReturnResponse(c *gin.Context, r global.Result, data any) {
	ReturnHttpResponse(c, http.StatusOK, r.Code(), r.Msg(), data)
}

// 成功业务码 + 数据
func ReturnSuccess(c *gin.Context, data any) {
	ReturnResponse(c, global.OkResult, data)
}

func ReturnError(c *gin.Context, r global.Result, data any) {
	slog.Info("[Func-ReturnError] " + r.Msg())

	var val string = r.Msg()

	if data != nil {
		switch v := data.(type) {
		case error:
			val = v.Error()
		case string:
			val = v
		}
		slog.Error(val) // 错误日志
	}

	// TODO: 2024/03/11 没搞懂这个函数的作用
	c.AbortWithStatusJSON(
		http.StatusOK,
		Response[any]{
			Code:    r.Code(),
			Message: r.Msg(),
			Data:    val,
		},
	)
}

// GetDB 从gin.Context中获取设置的*gorm.DB对象
func GetDB(c *gin.Context) *gorm.DB {
	return c.MustGet(global.CTX_DB).(*gorm.DB)
}

// GetRDB 获取建立好的*redis.Client
func GetRDB(c *gin.Context) *redis.Client {
	return c.MustGet(global.CTX_RDB).(*redis.Client)
}

func CurrentUserAuth(c *gin.Context) (*model.UserAuth, error) {
	key := global.CTX_USER_AUTH

	// 1
	if cache, exist := c.Get(key); exist && cache != nil {
		//slog.Debug("[Func-CurrentUserAuth] get from cache: " + cache.(*model.UserAuth).Username)
		return cache.(*model.UserAuth), nil
	}

	// 2
	session := sessions.Default(c)
	id := session.Get(key)
	if id == nil {
		return nil, errors.New("session 中没有 user_auth_id")
	}

	// 3
	db := GetDB(c)
	user, err := model.GetUserAuthInfoById(db, id.(int))
	if err != nil {
		return nil, err
	}

	c.Set(key, user)
	return user, nil
}
