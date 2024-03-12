package handle

import (
	"github.com/gin-gonic/gin"
	"my-gin-vue-blog/internal/global"
)

type UserAuth struct {
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (*UserAuth) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnError(c, global.ErrRequest, err)
		return
	}

	//db := GetDB(c)
}
