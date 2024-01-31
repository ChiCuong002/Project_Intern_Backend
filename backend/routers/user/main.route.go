package userRoute

import (
	"main/configs"
	"main/middleware"

	

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitUserRoute(app *echo.Echo) {
	router := app.Group("/user")
	router.Use(echojwt.WithConfig(configs.EchoJWTConfig()))
	router.Use(middleware.UserAuthentication)
	//profile 
	ProfileRouters(router)
	//product routers
	ProductManagementRoutes(router)
}
