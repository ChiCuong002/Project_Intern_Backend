package userRoute

import (
	controllers "main/handlers/controllers/product"

	"github.com/labstack/echo/v4"
)

func ProductManagementRoutes(router *echo.Group) {
	//add product
	router.POST("/add-product", controllers.AddProduct)
	//get product
	router.GET("/products", controllers.GetAllProduct)
	router.GET("/product/:id", controllers.DetailProduct)
	//update product
	router.PATCH("/update-product/:id", controllers.UpdateProduct)
}
