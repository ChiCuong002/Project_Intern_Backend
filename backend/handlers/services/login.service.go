package services

import (
	"errors"
	"fmt"
	storage "main/database"
	"main/helper/struct"
	"main/models"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// type jwtCustomClaims struct {
// 	UserId    uint   `json:"user_id"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	IsAdmin   uint   `json:"role_id"`
// 	jwt.RegisteredClaims
// }

func Login(username, password string) (models.User, string, error) {
	var user models.User
	db := storage.GetDB()
	//check email in db
	result := db.Where("phone_number = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, "", fmt.Errorf("User not registered: %v", result.Error)
	}
	//check match password
	if user.Password != password {
		return models.User{}, "", fmt.Errorf("Authentication error")
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
		return models.User{}, "", fmt.Errorf("Signed token error")
	}
	return user, t, nil
}
