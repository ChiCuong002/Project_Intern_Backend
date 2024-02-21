package services

import (
	storage "main/database"
	"main/schema"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ChangeInforProduct(c *schema.Product, db *gorm.DB, IDCategory int, ProductName string, Description string, Price float64, Quantity int, Image string) error {
	updates := map[string]interface{}{
		"category_id":  IDCategory,
		"product_name": ProductName,
		"description":  Description,
		"price":        Price,
		"quantity":     Quantity,
		"image_path":   Image,
	}
	result := db.Model(c).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func ChangeProduct(c echo.Context) error {
	var requestData struct {
		IDProduct   int     `json:"IDProduct"`
		IDCategory  int     `json:"IDCategory"`
		ProductName string  `json:"ProductName"`
		Description string  `json:"Description"`
		Price       float64 `json:"Price"`
		Quantity    int     `json:"NQuantity"`
		Image       string  `json:"Image"`
	}
	db := storage.GetDB()
	err := c.Bind(&requestData)
	if err != nil {
		//fmt.Println("Error binding request:", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "lấy san pham không được",
		})
	}
	var productToUpdate schema.Product
	result := db.First(&productToUpdate, requestData.IDProduct)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Không tìm thấy sản phẩm",
		})
	}
	err = ChangeInforProduct(&productToUpdate, db, requestData.IDCategory, requestData.ProductName, requestData.Description, requestData.Price, requestData.Quantity, requestData.Image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Lỗi thay đổi sản phẩm: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"product": requestData,
		"message": "Update sản phẩm thành công",
	})
}
