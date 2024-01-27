package services

import (
	"fmt"
	"main/schema"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func ChangePassword(u *schema.User,db *gorm.DB, newPassword, currentPassword string) error {
	hashNewPass, err := HashPassword(newPassword)
	fmt.Println("user: ", u)
	if err != nil {
		panic("ma hoa that bai")
	}
	match := CheckPasswordHash(currentPassword, u.Password)
	if !match {
		panic("mk ma hoa khong trung")
	}
	// Cập nhật mật khẩu người dùng trong cơ sở dữ liệu
	result := db.Model(u).Update("password", hashNewPass)
	if result.Error != nil {
		return result.Error
	}

	return nil
}