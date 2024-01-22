package controller

import (
	"main/schema"

	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {
	user := schema.User{}
	if err := c.Bind(&user); err != nil {
		
	}
}