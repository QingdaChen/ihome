package utils

import (
	"golang.org/x/crypto/bcrypt"
	"ihome/service/captcha/utils"
)

func Encryption(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.NewLog().Error("Encryption error")
	}
	return string(hash)
}

func CheckPasswd(password, passwdHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwdHash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
