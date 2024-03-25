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

// OptionVO
// TODO: 通用返回 用途? 2024/03/25 16:15
type OptionVO struct {
	ID   int    `json:"value"`
	Name string `json:"label"`
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

// Paginate 分页
func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

// Count 统计数量
func Count[T any](db *gorm.DB, data *T, where ...any) (int, error) {
	var total int64
	db = db.Model(data)
	if len(where) > 0 {
		db = db.Where(where[0], where[1:]...)
	}
	result := db.Count(&total)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(total), nil
}
