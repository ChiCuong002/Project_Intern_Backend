package adminRoute

import (
	controllers "main/handlers/controllers/product"

	"github.com/labstack/echo/v4"
)

func ProductManagementRoutes(router *echo.Group) {
	//all products
	router.GET("/products", controllers.GetAllProduct)
	//detail product
	router.GET("/product/:id", controllers.DetailProduct)
	//block product
	router.PATCH("/block-product/:id", controllers.DetailProduct)
}
