package adminRoute

import (
	controllers "main/handlers/controllers/user"

	"github.com/labstack/echo/v4"
)

func UserManagementRouters(router *echo.Group) {
	router.GET("/users/:id", controllers.DetailUser)                //user detail
	router.GET("/users", controllers.GetAllUser)                    //user list
	router.PATCH("/update-profile", controllers.UpdateUser)         //admin update admin's infomation
	router.POST("/changepassword", controllers.ChangePasswordUsers) //change user's password
	router.PATCH("/block/:id", controllers.BlockUser)               //block or unblock user
}
