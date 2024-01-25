package services

import (
	storage "main/database"
	"main/schema"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type InputData struct {
	Day   *int `json:"day"`
	Month int  `json:"month"`
}

func ThongKeTheoThangHandler(c echo.Context) error {
	var inputData InputData
	db := storage.GetDB()

	// Đọc dữ liệu từ request body
	if err := c.Bind(&inputData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	// Gọi hàm thống kê và trả về kết quả
	return ThongKeTheoThang(c, db, inputData)
}
func ThongKeTheoThang(c echo.Context, db *gorm.DB, input InputData) error {
	var total float64
	var orders []schema.Order
	if input.Month > 0 {
		if input.Day != nil && *input.Day > 0 {
			// Thống kê theo ngày và tháng
			if err := db.Where("EXTRACT(MONTH FROM order_date) = ? AND EXTRACT(DAY FROM order_date) = ?", input.Month, *input.Day).Find(&orders).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"error": "Internal server error",
				})
			}
		} else {
			// Thống kê theo tháng
			if err := db.Where("EXTRACT(MONTH FROM order_date) = ?", input.Month).Find(&orders).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"error": "Internal server error",
				})
			}
		}
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid input",
		})
	}

	for _, order := range orders {
		total += order.OrderTotal
	}

	// Trả về kết quả
	return c.JSON(http.StatusOK, echo.Map{
		"day":       input.Day,
		"month":     input.Month,
		"statistic": total,
	})
}
