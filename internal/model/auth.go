package model

import (
	"gorm.io/gorm"
	"time"
)

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
}

func GetUserAuthInfoByName(db *gorm.DB, name string) (*UserAuth, error) {
	var userAuth UserAuth
	result := db.Where(&UserAuth{Username: name}).First(&userAuth)
	return &userAuth, result.Error
}
