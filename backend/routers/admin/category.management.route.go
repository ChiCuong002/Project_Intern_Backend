package adminRoute

import (
	"main/handlers/controllers/category"

	"github.com/labstack/echo/v4"
)

func CategoriesManagementRouters(router *echo.Group) {
	router.GET("/categories", controllers.GetCategories)
	router.GET("/category/:id", controllers.DetailCategory)
}