package controllers

import (
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

func GetAllUser(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = PAGE_DEFAULT
		//return c.JSON(http.StatusBadRequest, "Invalid page parameter")
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		//return c.JSON(http.StatusBadRequest, "Invalid pageSize parameter")
		limit = LIMIT_DEFAULT
	}
	sort := c.QueryParam("sort")
	search := c.QueryParam("search")
	pagination := helper.Pagination{
		Page:   page,
		Limit:  limit,
		Sort:   sort,
		Search: search,
	}
	users, err := services.GetAllUserPagination(pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}
func DetailUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user, err := services.UserDetail(uint(idInt))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
func FindUser(c echo.Context) error {

	return nil
}
