package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Request().Header.Get("User-Role")

		if role == "admin" {
			log.Println("red button user detected")
		}

		return next(c)
	}
}
