package controllers

import (
	storage "main/database"
	"main/schema"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddCategoryRequest struct {
	NameCategory string `json:"category_name"`
}

func AddCategory(c echo.Context) error {
	var addCategoryRequest AddCategoryRequest
	db := storage.GetDB()
	var newUser schema.User
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	// Không cần kiểm tra loại đã tồn tại hay không

	newCategory := schema.Category{
		CategoryName: addCategoryRequest.NameCategory,
	}

	result := db.Create(&newCategory)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Lỗi: "+result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Thêm loại sản phẩm thành công!")
}
