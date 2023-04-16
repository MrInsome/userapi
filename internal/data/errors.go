package data

import (
	"github.com/go-chi/render"
	"github.com/labstack/echo/v4"
	"net/http"
)

var UserNotFound = map[string]string{
	"error": "user not found",
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) error {
	return echo.NewHTTPError(http.StatusBadRequest, ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	})
} //todo сделать кастомные ошибки (по желанию)
