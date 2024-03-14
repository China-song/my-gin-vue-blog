package utils

import "golang.org/x/crypto/bcrypt"

func BcryptHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// BcryptCheck 比较 明文字符串 和 哈希值
func BcryptCheck(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
