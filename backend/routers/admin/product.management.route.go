package adminRoute

import (
	controllers "main/handlers/controllers/product"

	"github.com/labstack/echo/v4"
)

func ProductManagementRoutes(router *echo.Group) {
	router.GET("/products", controllers.GetAllProduct)
}
