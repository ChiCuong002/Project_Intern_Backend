package controllers

import (
	"main/handlers/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ThongTinOrder struct {
	UserID    uint `json:"userID"`
	ProductID uint `json:"productID"`
	Quantity  uint `json:"quantity"`
}

func PurchaseProductController(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var thongtin ThongTinOrder

		// Bind JSON data from the request to thongtin
		if err := c.Bind(&thongtin); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
		}

		// Call the original PurchaseProduct function
		err := services.PurchaseProduct(db, c, thongtin.UserID, thongtin.ProductID, thongtin.Quantity)
		if err != nil {
			// Handle error if needed
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		// Return success response
		return c.JSON(http.StatusOK, echo.Map{"message": "Purchase successful"})
	}
}
