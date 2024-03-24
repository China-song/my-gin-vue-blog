package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAuth(t *testing.T) {
	db, err := initModelDB()
	assert.Nil(t, err)

	userAuth := UserAuth{
		Username: "12138",
		Password: "123456",
		UserInfo: &UserInfo{
			Nickname: "Leo",
		},
	}
	userAuth2 := UserAuth{
		Username: "12139",
		Password: "123456",
		UserInfo: &UserInfo{
			Nickname: "Mike",
		},
	}
	db.Create(&userAuth)
	db.Create(&userAuth2)
	user, err := GetUserAuthInfoById(db, userAuth.ID)
	assert.Nil(t, err)
	assert.Equal(t, userAuth.Username, user.Username)
	assert.Equal(t, userAuth.Password, user.Password)
	assert.Equal(t, userAuth.UserInfoId, user.UserInfoId)
	assert.Equal(t, userAuth.UserInfo.ID, user.UserInfo.ID)
	assert.Equal(t, userAuth.UserInfo.Nickname, user.UserInfo.Nickname)
	fmt.Println(userAuth.UserInfoId, user.UserInfoId)
	fmt.Println(userAuth.UserInfo.ID, user.UserInfo.ID)
	fmt.Println(userAuth2.UserInfo.ID, userAuth2.UserInfoId)
}
