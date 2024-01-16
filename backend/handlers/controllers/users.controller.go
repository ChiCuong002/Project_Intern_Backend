package controllers

import (
	"fmt"
	"main/handlers/services"
	helper "main/helper/struct"
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
		sortString = sortString + " asc"
	} else {
		sortString = sortString + " desc"
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
