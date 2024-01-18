package services

import (
	"errors"
	"fmt"
	storage "main/database"
	helper "main/helper/struct"
	"main/schema"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(username, password string) (schema.User, string, error) {
	var user schema.User
	db := storage.GetDB()
	//check phone number in db
	result := db.Where("phone_number = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return schema.User{}, "", fmt.Errorf("User not registered: %v", result.Error)
	}
	//check if user is active
	if !user.IsActive {
		return schema.User{}, "", fmt.Errorf("User is blocked")
	}
	//check match password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return schema.User{}, "", fmt.Errorf("Authentication error")
	}
	//create jwt token
	claims := &helper.JwtCustomClaims{
		UserId:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		IsAdmin:   user.RoleID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return schema.User{}, "", fmt.Errorf("Signed token error")
	}
	return user, t, nil
}
