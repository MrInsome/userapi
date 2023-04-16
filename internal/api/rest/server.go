package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func NewServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	//e.Use(middleware.RealIP()) перемещён в логгер
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 60 * time.Second,
	}))

	return e
}
