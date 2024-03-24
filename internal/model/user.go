package model

import "gorm.io/gorm"

type UserInfo struct {
	Model
	Email    string `json:"email" gorm:"type:varchar(30)"`
	Nickname string `json:"nickname" gorm:"unique;type:varchar(30);not null"`
	Avatar   string `json:"avatar" gorm:"type:varchar(1024);not null"`
	Intro    string `json:"intro" gorm:"type:varchar(255)"`
	Website  string `json:"website" gorm:"type:varchar(255)"`
}

func GetUserInfoById(db *gorm.DB, id int) (*UserInfo, error) {
	var userInfo UserInfo
	// result := db.Model(&userInfo).Where("id", id).First(&userInfo)
	result := db.Model(&userInfo).Where("id = ?", id).First(&userInfo)
	return &userInfo, result.Error
}
