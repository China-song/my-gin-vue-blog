package handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"my-gin-vue-blog/internal/global"
	"net/http"
)

// Response 响应结构体
type Response[T any] struct {
	Code    int    `json:"code"`    // 业务状态码
	Message string `json:"message"` // 响应消息
	Data    T      `json:"data"`    // 响应数据
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

// TODO: 返回gorm对象
//func GetDB(c *gin.Context) *gorm.DB {
//
//}
