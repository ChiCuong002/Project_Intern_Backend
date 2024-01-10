package main

import (
	storage "main/database"
	"main/handlers/controllers"
	helper "main/helper/struct"
	"main/middleware"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
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
func main() {
	e := echo.New()
	storage.InitDB()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/login", controllers.Login)

	//restricted group
	r := e.Group("/restricted")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))
	r.Use(middleware.AdminAuthentication)
	r.GET("", restricted)
	//users management
	//r.GET("/users", controllers.GetAllUser)
	r.GET("/users/:id", controllers.DetailUser)
	r.GET("/users/search", controllers.FindUser)
	r.GET("/users", controllers.GetAllUser)
	e.Logger.Fatal(e.Start(":8080"))
}
