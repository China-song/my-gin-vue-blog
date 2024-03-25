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
	categoryAPI handle.Category // 分类
	tagAPI      handle.Tag      // 标签
	articleAPI  handle.Article  // 文章
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

	// 用户模块
	//user := auth.Group("/user")
	//{
	//	user.GET("/list")
	//}

	// 分类模块
	category := auth.Group("/category")
	{
		category.GET("/list", categoryAPI.GetList)     // 分类列表
		category.POST("", categoryAPI.SaveOrUpdate)    // 新增/编辑分类
		category.DELETE("", categoryAPI.Delete)        // 删除分类
		category.GET("/option", categoryAPI.GetOption) // 分类选项列表
	}

	//标签模块
	tag := auth.Group("/tag")
	{
		tag.GET("/list", tagAPI.GetList)     // 标签列表
		tag.POST("", tagAPI.SaveOrUpdate)    // 新增/编辑标签
		tag.DELETE("", tagAPI.Delete)        // 删除标签
		tag.GET("/option", tagAPI.GetOption) // 标签选项列表
	}

	// 文章模块
	articles := auth.Group("/article")
	{
		articles.GET("/list", articleAPI.GetList)  // 文章列表
		articles.POST("", articleAPI.SaveOrUpdate) // 新增/编辑文章
		articles.PUT("/top", articleAPI.UpdateTop) // 更新文章置顶
		articles.GET("/:id", articleAPI.GetDetail) // 文章详情
		articles.PUT("/soft-delete", articleAPI.UpdateSoftDelete)
		articles.DELETE("", articleAPI.Delete) // 物理删除文章
	}
}
