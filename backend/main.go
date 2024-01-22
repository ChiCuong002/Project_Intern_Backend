package main

import (
	storage "main/database"
	"main/handlers/controllers"
	productController "main/handlers/controllers/product"
	userController "main/handlers/controllers/user"
	helper "main/helper/struct"
	"main/middleware"
	"main/schema"
	"net/http"

	//"main/schema"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	gomiddleware "github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	UserId    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsAdmin   uint   `json:"role_id"`
	jwt.RegisteredClaims
}

func restricted(c echo.Context) error {
	return c.JSON(http.StatusOK, "Admin here")
}
func api(c echo.Context) error {
	return c.JSON(http.StatusOK, "User here")
}
func main() {
	e := echo.New()
	schema.Migration()
	e.Use(gomiddleware.Logger())
	storage.InitDB()
	//echojwt config
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	//CORS config for all routes
	CORSConfig := gomiddleware.CORSWithConfig(gomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	})
	//Apply CORS middleware to all routes
	e.Use(CORSConfig)

	//main route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//register route
	e.POST("/register", controllers.RegisterUser)
	//login route
	e.POST("/login", controllers.Login)
	//categories drop down
	e.GET("/categories-dropdown", controllers.CategoriesDropDown)
	//user route
	u := e.Group("/api")
	u.Use(echojwt.WithConfig(config))
	u.Use(middleware.UserAuthentication)
	u.GET("", api)
	u.POST("/addproduct", productController.AddProduct)
	//restricted group
	r := e.Group("/restricted")

	//auth middleware
	r.Use(echojwt.WithConfig(config))
	r.Use(middleware.AdminAuthentication)

	//restricted

	r.GET("", restricted)
	//users management
	r.GET("/users", controllers.GetAllUser)                    //user list
	r.GET("/users/:id", controllers.DetailUser)                //user detail
	r.POST("/update-profile", userController.UpdateUser)       //admin update admin's infomation
	r.POST("/changepassword", controllers.ChangePasswordUsers) //change user's password
	r.PATCH("/block/:id", controllers.BlockUser)               //block or unblock user
	//categories management
	r.GET("/categories", controllers.GetCategories)
	r.GET("/category/:id", controllers.DetailCategory)
	e.Logger.Fatal(e.Start(":8080"))
}
