package handle

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"my-gin-vue-blog/internal/global"
	"my-gin-vue-blog/internal/model"
	"my-gin-vue-blog/internal/utils"
	"my-gin-vue-blog/internal/utils/jwt"
	"strconv"
)

type UserAuth struct {
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginVO 用户登录成功返回用户信息和token
type LoginVO struct {
	model.UserInfo

	// 点赞 Set: 用于记录用户点赞过的文章, 评论
	ArticleLikeSet []string `json:"article_like_set"`
	CommentLikeSet []string `json:"comment_like_set"`
	Token          string   `json:"token"`
}

func (*UserAuth) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnError(c, global.ErrRequest, err)
		return
	}

	db := GetDB(c)
	rdb := GetRDB(c)

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

	// 检查密码是否正确
	if !utils.BcryptCheck(userAuth.Password, req.Password) {
		ReturnError(c, global.ErrPassword, nil)
		return
	}

	// 登录成功后获取该用户的信息：用户信息 用户权限
	userInfo, err := model.GetUserInfoById(db, userAuth.UserInfoId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ReturnError(c, global.ErrUserNotExist, nil)
			return
		}
		ReturnError(c, global.ErrDbOp, err)
		return
	}

	roleIds, err := model.GetRoleIdsByUserId(db, userAuth.ID)
	if err != nil {
		ReturnError(c, global.ErrDbOp, err)
		return
	}

	// redis 获取文章点赞 评论点赞
	articleLikeSet, err := rdb.SMembers(rctx, global.ARTICLE_USER_LIKE_SET+strconv.Itoa(userAuth.ID)).Result()
	if err != nil {
		ReturnError(c, global.ErrRedisOp, err)
		return
	}

	commentLikeSet, err := rdb.SMembers(rctx, global.COMMENT_USER_LIKE_SET+strconv.Itoa(userAuth.ID)).Result()
	if err != nil {
		ReturnError(c, global.ErrRedisOp, err)
		return
	}

	// 登录成功 生成带有用户信息的Token
	jwtConf := global.Conf.JWT
	token, err := jwt.GenToken(jwtConf.Secret, jwtConf.Issuer, int(jwtConf.Expire), userAuth.ID, roleIds)
	if err != nil {
		ReturnError(c, global.ErrTokenCreate, err)
		return
	}

	// TODO: 更新ip地址 登录时间

	//slog.Info("用户登录成功: " + userAuth.Username)
	log.Println("用户登录成功: ", userAuth.Username)

	// TODO: session? 2024/03/21
	session := sessions.Default(c)
	session.Set(global.CTX_USER_AUTH, userAuth.ID)
	session.Save()

	// 删除 Redis 中的离线状态
	_, err = rdb.Del(rctx, global.OFFLINE_USER+strconv.Itoa(userAuth.ID)).Result()
	if err != nil {
		ReturnError(c, global.ErrRedisOp, err)
		return
	}

	ReturnSuccess(c, LoginVO{
		UserInfo:       *userInfo,
		ArticleLikeSet: articleLikeSet,
		CommentLikeSet: commentLikeSet,
		Token:          token,
	})
}
