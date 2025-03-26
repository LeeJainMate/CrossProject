package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// ตรวจสอบรหัสผ่าน
func CheckPassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
