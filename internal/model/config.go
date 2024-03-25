package model

import "gorm.io/gorm"

type Config struct {
	Model
	Key   string `gorm:"unique;type:varchar(256)" json:"key"`
	Value string `gorm:"type:varchar(256)" json:"value"`
	Desc  string `gorm:"type:varchar(256)" json:"desc"`
}

func GetConfig(db *gorm.DB, key string) string {
	var config Config
	result := db.Where("key", key).First(&config)

	if result.Error != nil {
		return ""
	}

	return config.Value
}
