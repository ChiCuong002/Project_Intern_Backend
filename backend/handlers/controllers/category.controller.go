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
	var newCategory schema.Category
	err := c.Bind(&newCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	// Không cần kiểm tra loại đã tồn tại hay không

	newCategory = schema.Category{
		CategoryName: addCategoryRequest.NameCategory,
	}

	result := db.Create(&newCategory)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Lỗi: "+result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Thêm loại sản phẩm thành công!")
}

func EditCategory(c echo.Context) error {
	var requestData struct {
		IDCategory      string `json:IDCategory`
		NewNameCategory string `json:"NewNameCategory"`
	}
	db := storage.GetDB()
	err := c.Bind(&requestData)
	if err != nil {
		//fmt.Println("Error binding request:", err)
		return c.JSON(http.StatusBadRequest, "lấy loại không được")
	}

	var categoyToUpdate schema.Category
	result := db.First(&categoyToUpdate, requestData.IDCategory)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Không tìm thấy loại sản phẩm")
	}

	err = categoyToUpdate.ChangeNameCategory(db, requestData.NewNameCategory)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Lỗi thay đổi loại: "+err.Error())
	}

	return c.JSON(http.StatusOK, "Đổi tên loại thành công")
}
