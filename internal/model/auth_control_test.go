package model

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

func initModelDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //
		},
	})
	if err != nil {
		return nil, err
	}

	err = MakeMigrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestRoleAndResource(t *testing.T) {
	db, err := initModelDB()
	assert.Nil(t, err)
	// 添加两个权限
	p1, err := AddResource(db, "api_1", "/v1/api_1", "GET", false)
	assert.Nil(t, err)
	p2, err := AddResource(db, "api_2", "/v1/api_2", "POST", false)
	assert.Nil(t, err)
	// 创建一个角色带有这俩权限
	role1, err := AddRoleWithResources(db, "admin", "管理员", []int{p1.ID, p2.ID})
	assert.Nil(t, err)
	// 获取找个个角色的权限
	resources, err := GetResourcesByRole(db, role1.ID)
	assert.Nil(t, err)
	assert.Len(t, resources, 2)
	flag, err := CheckRoleAuth(db, role1.ID, p1.Url, p1.Method)
	assert.Nil(t, err)
	assert.True(t, flag)
}
