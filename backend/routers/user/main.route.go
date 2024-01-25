package userRoute

import (
	"main/configs"
	"main/middleware"

	productController "main/handlers/controllers/product"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitUserRoute(app *echo.Echo) {
	router := app.Group("/user")
	router.Use(echojwt.WithConfig(configs.EchoJWTConfig()))
	router.Use(middleware.UserAuthentication)
	//add product
	router.POST("/add-product", productController.AddProduct)
	//get product
	router.GET("/product/:id", productController.DetailProduct)
}
