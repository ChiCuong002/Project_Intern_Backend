package controllers

import (
	//"crypto/ecdsa"
	"fmt"
	"main/handlers/services"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
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
