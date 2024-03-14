package handle

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my-gin-vue-blog/internal/global"
	"my-gin-vue-blog/internal/model"
	"my-gin-vue-blog/internal/utils"
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

	db := GetDB(c)
	userAuth, err := model.GetUserAuthInfoByName(db, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ReturnError(c, global.ErrUserNotExist, nil)
			return
		}
		// TODO: 什么情况下返回的data是error 什么情况下是对象
		ReturnError(c, global.ErrDbOp, err)
		return
	}

	// TODO: 接下来检查密码 加密
	if !utils.BcryptCheck(userAuth.Password, req.Password) {
		ReturnError(c, global.ErrPassword, nil)
		return
	}
}
