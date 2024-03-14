package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBcrypt(t *testing.T) {
	password := "123456"

	// 加密
	hashedPassword, err := BcryptHash(password)
	assert.Nil(t, err)

	// 验证
	result := BcryptCheck(hashedPassword, password)
	assert.True(t, result)
}
