package global

// Redis Key
const (
	OFFLINE_USER = "offline_user:" // 强制下线用户

	ARTICLE_USER_LIKE_SET = "article_user_like:" // 文章点赞 Set
	ARTICLE_LIKE_COUNT    = "article_like_count" // 文章点赞数
	ARTICLE_VIEW_COUNT    = "article_view_count" // 文章查看数

	COMMENT_USER_LIKE_SET = "comment_user_like:" // 评论点赞 Set
)

// Gin Context Key
const (
	CTX_DB        = "_db_field"
	CTX_RDB       = "_rdb_field"
	CTX_USER_AUTH = "_user_auth_field"
)

// Config Key
const (
	CONFIG_ARTICLE_COVER = "article_cover"
)
