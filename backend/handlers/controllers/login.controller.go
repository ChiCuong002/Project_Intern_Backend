package controllers

import (
	//"crypto/ecdsa"
	"fmt"
	storage "main/database"
	"main/handlers/services"
	"main/models"
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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
	}
	user := models.User{}
	username := LoginRequest.PhoneNumber
	password := LoginRequest.Password
	fmt.Println("Received phone_number:", username)
	fmt.Println("Received password:", password)
	//get user token error
	user, token, err := services.Login(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"user":  user,
		"token": token,
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
			"message": "khong lay duoc du lieu",
		})
	}
	// Kiểm tra xem email đã tồn tại chưa
	var existingUser schema.User
	result := db.Where("phone_number = ?", newUser.PhoneNumber).First(&existingUser)
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusConflict, echo.Map{
			"message": "Số điện thoại đã được đăng ký",
		})
	}
	if len(newUser.Password) < 8 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Mật khẩu phải trên 8 ký tự",
		})
	}
	hash, err := HashPassword(newUser.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Lỗi khi mã hóa password",
		})
	}
	match := CheckPasswordHash(newUser.Password, hash)
	if !match {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Lỗi khi check password",
		})
	}
	newUser.Password = hash
	if newUser.RoleID == 0 {
		newUser.RoleID = 2
	}
	result = db.Create(&newUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user":    newUser,
		"message": "Register succesfull",
	})
}
