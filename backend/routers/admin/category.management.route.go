package adminRoute

import (
	controllers "main/handlers/controllers/category"

	"github.com/labstack/echo/v4"
)

func CategoriesManagementRouters(router *echo.Group) {
	router.GET("/categories", controllers.GetCategories)
	router.GET("/category/:id", controllers.DetailCategory)
	router.POST("/add-category", controllers.AddCategory)
	router.PATCH("/update-category", controllers.EditCategory)
}
