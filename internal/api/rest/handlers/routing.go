package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"refactoring/internal/api/rest"
	"time"
)

func NewRouter() *echo.Echo {
	e := rest.NewServer()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().String())
	})

	e.GET("/api/v1/users", searchUsers)
	e.POST("/api/v1/users", createUser)
	e.GET("/api/v1/users/:id", getUser)
	e.PATCH("/api/v1/users/:id", updateUser)
	e.DELETE("/api/v1/users/:id", deleteUser)

	return e
}
