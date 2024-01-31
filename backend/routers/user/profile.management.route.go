package userRoute

import (
	controllers "main/handlers/controllers/user"

	"github.com/labstack/echo/v4"
)

func ProfileRouters(router *echo.Group) {
	router.PATCH("/update-profile", controllers.UpdateUser)
}