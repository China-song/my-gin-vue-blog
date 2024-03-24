package model

import (
	"gorm.io/gorm"
	"time"
)

// UserAuth 权限相关 登录检查
// TODO: 相关字段不清晰 2024/03/18 23:06
type UserAuth struct {
	Model
	Username      string     `json:"username" gorm:"unique;type:varchar(50)"`
	Password      string     `json:"-" gorm:"type:varchar(100)"` // -表示不发密码？
	LoginType     int        `json:"login_type" gorm:"type:tinyint(1);comment:登录类型"`
	IpAddress     string     `json:"ip_address" gorm:"type:varchar(50);comment:登录IP地址"`
	IpSource      string     `json:"ip_source" gorm:"type:varchar(50);comment:IP来源"`
	LastLoginTime *time.Time `json:"last_login_time"` // 指针的意思？
	IsDisable     bool       `json:"is_disable"`
	IsSuper       bool       `json:"is_super"`

	UserInfoId int       `json:"user_info_id"`
	UserInfo   *UserInfo `json:"info"`                                  // info -> user_info?
	Roles      []*Role   `json:"roles" gorm:"many2many:user_auth_role"` // 为什么这里是指针
}

type Role struct {
	Model
	Name      string `gorm:"unique" json:"name"`
	Label     string `gorm:"unique" json:"label"`
	IsDisable bool   `json:"is_disable"`

	Resources []Resource `json:"resources" gorm:"many2many:role_resource"`
	Menus     []Menu     `json:"menus" gorm:"many2many:role_menu"`
	Users     []UserAuth `json:"users" gorm:"many2many:user_auth_role"`
}

type Resource struct {
	Model
	Name      string `gorm:"unique;type:varchar(50)" json:"name"`
	ParentId  int    `json:"parent_id"`
	Url       string `gorm:"type:varchar(255)" json:"url"`
	Method    string `gorm:"type:varchar(10)" json:"request_method"`
	Anonymous bool   `json:"is_anonymous"`

	Roles []*Role `json:"roles" gorm:"many2many:role_resource"`
}

type Menu struct {
	Model
	ParentId     int    `json:"parent_id"`
	Name         string `gorm:"uniqueIndex:idx_name_and_path;type:varchar(20)" json:"name"` // 菜单名称
	Path         string `gorm:"uniqueIndex:idx_name_and_path;type:varchar(50)" json:"path"` // 路由地址
	Component    string `gorm:"type:varchar(50)" json:"component"`                          // 组件路径
	Icon         string `gorm:"type:varchar(50)" json:"icon"`                               // 图标
	OrderNum     int8   `json:"order_num"`                                                  // 排序
	Redirect     string `gorm:"type:varchar(50)" json:"redirect"`                           // 重定向地址
	Catalogue    bool   `json:"is_catalogue"`                                               // 是否为目录
	Hidden       bool   `json:"is_hidden"`                                                  // 是否隐藏
	KeepAlive    bool   `json:"keep_alive"`                                                 // 是否缓存
	External     bool   `json:"is_external"`                                                // 是否外链
	ExternalLink string `gorm:"type:varchar(255)" json:"external_link"`                     // 外链地址

	Roles []*Role `json:"roles" gorm:"many2many:role_menu"`
}

// 关联表

type RoleResource struct {
	RoleId     int `json:"-" gorm:"primaryKey;uniqueIndex:idx_role_resource"` // 复合主键
	ResourceId int `json:"-" gorm:"primaryKey;uniqueIndex:idx_role_resource"`
}

type UserAuthRole struct {
	UserAuthId int `gorm:"primaryKey;uniqueIndex:idx_user_auth_role"`
	RoleId     int `gorm:"primaryKey;uniqueIndex:idx_user_auth_role"`
}

type RoleMenu struct {
	RoleId int `json:"-" gorm:"primaryKey;uniqueIndex:idx_role_menu"`
	MenuId int `json:"-" gorm:"primaryKey;uniqueIndex:idx_role_menu"`
}

func GetUserAuthInfoByName(db *gorm.DB, name string) (*UserAuth, error) {
	var userAuth UserAuth
	result := db.Where(&UserAuth{Username: name}).First(&userAuth)
	return &userAuth, result.Error
}

func GetRoleIdsByUserId(db *gorm.DB, userAuthId int) (ids []int, err error) {
	result := db.Model(&UserAuthRole{UserAuthId: userAuthId}).Pluck("role_id", &ids)
	return ids, result.Error
}

// UserAuth

// GetUserAuthInfoById 返回UserAuth 及用户信息 及用户角色
func GetUserAuthInfoById(db *gorm.DB, id int) (*UserAuth, error) {
	var userAuth = UserAuth{Model: Model{ID: id}}
	result := db.Model(&userAuth).
		Preload("Roles").Preload("UserInfo").
		First(&userAuth)
	return &userAuth, result.Error
}
