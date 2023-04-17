package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"refactoring/internal/api/rest"
	"refactoring/internal/config"
	"refactoring/internal/services"
	"time"
)

type Routing struct {
	echoRouter *echo.Echo
	service    *services.Service
	config     *config.Configs
}

func NewRouter(service *services.Service, config *config.Configs) *Routing {
	return &Routing{service: service, config: config}
}

func (r *Routing) Start() *echo.Echo {
	e := rest.NewServer()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().String())
	})

	e.GET("/api/v1/users", r.SearchUsers)
	e.POST("/api/v1/users", r.CreateUser)
	e.GET("/api/v1/users/:id", r.GetUser)
	e.PATCH("/api/v1/users/:id", r.UpdateUser)
	e.DELETE("/api/v1/users/:id", r.DeleteUser)

	return e
}
