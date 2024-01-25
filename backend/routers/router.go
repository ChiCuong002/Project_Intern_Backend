package routers

import (
	adminRoute "main/routers/admin"
	normalRoute "main/routers/normal"
	userRoute "main/routers/user"

	"github.com/labstack/echo/v4"
)

func Routers(app *echo.Echo) {
	normalRoute.InitNormalRouters(app)
	adminRoute.InitAdminRoute(app)
	userRoute.InitUserRoute(app)
}
