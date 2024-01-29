package middleware

import (
	"fmt"
	helper "main/helper/struct"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	ADMIN = 1
	USER  = 2
)

func AdminAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusNotFound, "Error: JWT token missing or invalid")
		}
		fmt.Println("token: ", user)
		claims := user.Claims.(*helper.JwtCustomClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusNotFound, "Error: Failed to cast claims as jwtCustomClaims")
		}
		fmt.Println("claims: ", claims)
		fmt.Println("isAdmin: ", claims.IsAdmin)
		if claims.IsAdmin != ADMIN {
			return echo.NewHTTPError(http.StatusUnauthorized, "Error: Authorization Required")
		}
		c.Set("userID", claims.UserId)
		return next(c)
	}
}
func UserAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusNotFound, "Error: JWT token missing or invalid")
		}
		fmt.Println("token: ", user)
		claims := user.Claims.(*helper.JwtCustomClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusNotFound, "Error: Failed to cast claims as jwtCustomClaims")
		}
		fmt.Println("claims: ", claims)
		fmt.Println("isAdmin: ", claims.IsAdmin)
		if claims.IsAdmin != USER {
			return echo.NewHTTPError(http.StatusUnauthorized, "Error: Authorization Required")
		}
		c.Set("userID", claims.UserId)
		return next(c)
	}
}
