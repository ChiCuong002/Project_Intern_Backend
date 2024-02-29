package userRoute

import (
	controllers "main/handlers/controllers/product"

	"github.com/labstack/echo/v4"
)

func ProductManagementRoutes(router *echo.Group) {
	//add product
	router.POST("/add-product", controllers.AddProduct)
	//get product inventory
	router.GET("/my-inventory", controllers.MyInventory)
	//get posted product
	router.GET("/my-products", controllers.MyProduct)
	//detail product
	router.GET("/product/:id", controllers.DetailProduct)
	//update product
	router.PATCH("/update-product/:id", controllers.UpdateProduct)
	//block product
	router.PATCH("/block-product/:id", controllers.BlockProduct)
}
