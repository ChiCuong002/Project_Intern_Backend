package controllers

import (
	"fmt"
	storage "main/database"
	"main/handlers/services"
	helper "main/helper/struct"
	"main/schema"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	LIMIT_DEFAULT = 10
	PAGE_DEFAULT  = 1
)

func sortString(sort string) string {
	order := sort[0]
	sortString := sort[0:]
	if rune(order) == '+' {
		sortString = sortString + " desc"
	} else {
		sortString = sortString + " asc"
	}
	fmt.Println("sortString: ", sortString)
	return sortString
}
func GetAllUser(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = PAGE_DEFAULT
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = LIMIT_DEFAULT
	}
	sort := c.QueryParam("sort")
	if sort != "" {
		sort = sortString(sort)
	}
	search := c.QueryParam("search")
	pagination := helper.Pagination{
		Page:   page,
		Limit:  limit,
		Sort:   sort,
		Search: search,
	}
	users, err := services.GetAllUserPagination(pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}
func DetailUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	fmt.Println("id: ", idInt)
	user, err := services.UserDetail(uint(idInt))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("user: ", user)
	return c.JSON(http.StatusOK, user)
}

func ChangePasswordUsers(c echo.Context) error {
	db := storage.GetDB()
	var requestData struct {
		UserID          uint   `json:"UserID"`
		CurrentPassword string `json:"CurrentPassword"`
		NewPassword     string `json:"NewPassword"`
	}
	err := c.Bind(&requestData)
	if err != nil {
		fmt.Println("Error binding request:", err)
		return c.JSON(http.StatusBadRequest, "lấy người dùng ko đc")
	}

	fmt.Println("User ID:", requestData.UserID)
	fmt.Println("New Password:", requestData.NewPassword)

	var userToUpdate schema.User
	result := db.First(&userToUpdate, requestData.UserID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Không tìm thấy người dùng")
	}

	// Thay đổi mật khẩu của người dùng
	err = userToUpdate.ChangePassword(db, requestData.NewPassword, requestData.CurrentPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Lỗi thay đổi mật khẩu: "+err.Error())
	}

	return c.JSON(http.StatusOK, "Thay đổi mật khẩu thành công")
}
func BlockUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	fmt.Println("id: ", idInt)
	user, err := services.BlockUser(uint(idInt))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Block/Unblock user successfully",
		"user":    user,
	})
}
