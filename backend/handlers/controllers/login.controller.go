package controllers

import (
	//"crypto/ecdsa"
	"fmt"
	storage "main/database"
	"main/handlers/services"
	//"main/helper/validation"
	"main/schema"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func Login(c echo.Context) error {
	var LoginRequest LoginRequest
	if err := c.Bind(&LoginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid JSON format"})
	}
	// //phone number validate
	// if !validation.IsPhoneNumber(LoginRequest.PhoneNumber) {
	// 	return c.JSON(http.StatusBadRequest, echo.Map{
	// 		"error": "Invalid phone number format",
	// 	})
	// }
	// //password validate
	// isPassword, errs := validation.IsPassword(LoginRequest.Password)
	// if !isPassword {
	// 	return c.JSON(http.StatusBadRequest, echo.Map{
	// 		"errors": errs,
	// 	})
	// }
	//
	user := schema.User{}
	username := LoginRequest.PhoneNumber
	password := LoginRequest.Password
	fmt.Println("Received phone_number:", username)
	fmt.Println("Received password:", password)
	//get user token error
	user, token, err := services.Login(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message" : err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"user":  user,
		"token": token,
		"message": "User login successfull",
	})
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func RegisterUser(c echo.Context) error {
	// Đọc dữ liệu từ request
	db := storage.GetDB()
	var newUser schema.User
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message" : "Invalid request payload",
		})
	}
	// Kiểm tra xem email đã tồn tại chưa
	var existingUser schema.User
	result := db.Where("phone_number = ?", newUser.PhoneNumber).First(&existingUser)
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusConflict, echo.Map{
			"message" : "Phone number is already exists",
		})
	}
	if len(newUser.Password) < 8 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message" : "Password should have at least 8 characters",
		})
	}
	hash, err := HashPassword(newUser.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message" : "Hash password failed", 
		})
	}
	match := CheckPasswordHash(newUser.Password, hash)
	if !match {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message" : "Compare password failed",
		})
	}
	newUser.Password = hash
	if newUser.RoleID == 0 {
		newUser.RoleID = 2
	}
	result = db.Create(&newUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message" : result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message" : "Register successfull",
	})
}
