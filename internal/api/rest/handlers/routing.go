package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"refactoring/internal/api/rest"
	"refactoring/internal/data"
	"time"
)

func NewRouter() *echo.Echo {
	e := rest.NewServer()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().String())
	})

	e.GET("/api/v1/users", data.SearchUsers)
	e.POST("/api/v1/users", data.CreateUser)
	e.GET("/api/v1/users/:id", data.GetUser)
	e.PATCH("/api/v1/users/:id", data.UpdateUser)
	e.DELETE("/api/v1/users/:id", data.DeleteUser)

	return e
}
