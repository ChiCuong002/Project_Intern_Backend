package userRoute

import (
	controllers "main/handlers/controllers/user"

	"github.com/labstack/echo/v4"
)

func ProfileRouters(router *echo.Group) {
	//get user profile
	router.GET("/my-profile", controllers.MyProfile)
	//update user profile
	router.PATCH("/update-profile", controllers.UpdateUser)
}
