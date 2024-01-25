package userRoute

import (
	"main/configs"
	"main/middleware"

	controllers "main/handlers/controllers/product"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitUserRoute(app *echo.Echo) {
	router := app.Group("/user")
	router.Use(echojwt.WithConfig(configs.EchoJWTConfig()))
	router.Use(middleware.UserAuthentication)
	//add product
	router.POST("/add-product", controllers.AddProduct)
	//get product
	router.GET("/products", controllers.GetAllProduct)
	router.GET("/product/:id", controllers.DetailProduct)
	//update product
	router.PATCH("/update-product/:id", controllers.UpdateProduct)
}
