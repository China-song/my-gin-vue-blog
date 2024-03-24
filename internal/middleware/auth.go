package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my-gin-vue-blog/internal/global"
	"my-gin-vue-blog/internal/handle"
	"my-gin-vue-blog/internal/model"
)

// JWTAuth 首先查看请求是否需要权限 不需要就不用看token
// 否则 对token：获取、解析、查看是否过期、
// 猜测：如果token包含的用户信息具有的权限能够执行请求 则处理后续请求 否则返回
// TODO: 获取其包含的用户信息 保存到session中？ 2024/03/21 15:27
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// PermissionCheck 根据用户角色所具有的权限 检查 用户请求是否合法
func PermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetBool("skip_check") {
			c.Next()
			return
		}

		// 获取用户权限
		db := c.MustGet(global.CTX_DB).(*gorm.DB)
		auth, err := handle.CurrentUserAuth(c)
		if err != nil {
			handle.ReturnError(c, global.ErrUserNotExist, err)
			return
		}

		// 具有最高权限
		if auth.IsSuper {
			c.Next()
			return
		}

		url := c.FullPath()[4:]
		method := c.Request.Method
		// 查找这条权限的ResourceID

		// 查看
		for _, role := range auth.Roles {
			// 一个role具有好多权限
			// 查看找个role
			pass, err := model.CheckRoleAuth(db, role.ID, url, method)
			if err != nil {
				handle.ReturnError(c, global.ErrDbOp, err)
			}
			if pass {
				c.Next()
				return
			}
		}
		handle.ReturnError(c, global.ErrPermission, nil)
		return
	}
}
