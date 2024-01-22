package controller

import (
	//"main/schema"
	"fmt"
	service "main/handlers/services/user"
	helper "main/helper/struct"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {
	fmt.Println("Updated user")
	userID := c.Get("userID").(uint)
	data := helper.UpdateData{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Error binding data",
		})
	}
	data.UserID = userID
	err := service.UpdateUser(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User infomation is updated successfully",
	})
}
