package controller

import (
	//"main/schema"
	"fmt"
	storage "main/database"
	service "main/handlers/services/user"
	userServices "main/handlers/services/user"
	helper "main/helper/struct"
	"main/schema"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	LIMIT_DEFAULT = 10
	PAGE_DEFAULT  = 1
	SORT_DEFAULT  = " user_id desc"
)

func sortString(sort string) string {
	order := sort[0]
	sortString := sort[0:]
	if rune(order) == '+' {
		sortString = sortString + " desc"
	} else if rune(order) == '-' {
		sortString = sortString + " asc"
	} else {
		sortString = ""
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
	} else {
		sort = SORT_DEFAULT
	}
	search := c.QueryParam("search")
	pagination := helper.Pagination{
		Page:   page,
		Limit:  limit,
		Sort:   sort,
		Search: search,
	}
	users, err := userServices.GetAllUserPagination(pagination)
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
	user, err := userServices.UserDetail(uint(idInt))
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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Binding data failed",
		})
	}

	fmt.Println("User ID:", requestData.UserID)
	fmt.Println("New Password:", requestData.NewPassword)

	var userToUpdate schema.User
	result := db.First(&userToUpdate, requestData.UserID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	// Thay đổi mật khẩu của người dùng
	err = userToUpdate.ChangePassword(db, requestData.NewPassword, requestData.CurrentPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Password is changed",
	})
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
	user, err := userServices.BlockUser(uint(idInt))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Block/Unblock user successfully",
		"user":    user,
	})
}

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
