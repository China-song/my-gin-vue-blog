package model

import (
	"gorm.io/gorm"
	"time"
)

// 通用模型

type Model struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MakeMigrate 迁移数据库表
func MakeMigrate(db *gorm.DB) error {
	// 设置关联表
	db.SetupJoinTable(&Role{}, "Menus", &RoleMenu{})
	db.SetupJoinTable(&Role{}, "Resources", &RoleResource{})
	db.SetupJoinTable(&Role{}, "Users", &UserAuthRole{})
	db.SetupJoinTable(&UserAuth{}, "Roles", &UserAuthRole{})

	return db.AutoMigrate(
		&UserInfo{}, // 用户信息

		&UserAuth{},     // 用户验证
		&Role{},         // 角色
		&Menu{},         // 菜单
		&Resource{},     // 资源（接口）
		&RoleMenu{},     // 角色-菜单 关联
		&RoleResource{}, // 角色-资源 关联
		&UserAuthRole{}, // 用户-角色 关联
	)
}
