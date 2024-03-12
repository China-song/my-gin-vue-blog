package model

import (
	"gorm.io/gorm"
	"time"
)

// 通用模型

type Model struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MakeMigrate 迁移数据库表
func MakeMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&UserAuth{}, // 用户验证
	)
}
