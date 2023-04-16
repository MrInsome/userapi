package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"refactoring/internal/data"
)

func (r *Routing) SearchUsers(c echo.Context) error {
	f, err := r.service.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, f)
}

func (r *Routing) GetUser(c echo.Context) error {
	id := c.Param("id")
	f, err := r.service.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, data.ErrInvalidRequest(err))
	} else {
		return c.JSON(http.StatusOK, f)
	}
}

func (r *Routing) CreateUser(c echo.Context) error {
	var user data.User
	request := data.CreateUserRequest{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	user.Email = request.Email
	user.DisplayName = request.DisplayName
	userid, err := r.service.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, data.ErrInvalidRequest(err))
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"user_id": userid,
	})
}

func (r *Routing) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user, err := r.service.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data.ErrInvalidRequest(err))
	}
	request := data.UpdateUserRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	oldName := user.DisplayName
	user.DisplayName = request.DisplayName
	err = r.service.UpdateUser(id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data.ErrInvalidRequest(err))
	}
	return c.JSON(http.StatusOK, "Имя изменено c '"+oldName+"' на '"+request.DisplayName+"'")
}

func (r *Routing) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := r.service.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Пользователь с id:"+id+" удалён.")
}
