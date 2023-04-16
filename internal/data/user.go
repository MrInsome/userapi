package data

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"time"
)

const Store = `users.json`
const StoreDir = "."

func SearchUsers(c echo.Context) error {
	f, err := os.ReadFile(Store)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	s := UserStore{}
	err = json.Unmarshal(f, &s)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, s.List)
}

func CreateUser(c echo.Context) error {
	var s UserStore
	StoreFS := http.Dir(StoreDir)
	_, err := StoreFS.Open(Store)
	if err == nil {
		f, err := os.ReadFile(Store)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		err = json.Unmarshal(f, &s)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	} else if !errors.Is(err, fs.ErrNotExist) {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	request := CreateUserRequest{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	s.Increment = 0
	for _, user := range s.List {
		if user.Email == request.Email {
			return echo.NewHTTPError(http.StatusBadRequest, "Пользователь с таким email уже существует.")
		}
	}
	for _, user := range s.List {
		if user.ID > s.Increment {
			s.Increment = user.ID - 1
			break
		}
		s.Increment++
	}
	u := User{
		ID:          s.Increment,
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = u

	b, err := json.Marshal(&s)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = os.WriteFile(Store, b, fs.ModePerm)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"user_id": id,
	})
}

func GetUser(c echo.Context) error {
	f, err := os.ReadFile(Store)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	s := UserStore{}
	_ = json.Unmarshal(f, &s)

	id := c.Param("id")

	if user, ok := s.List[id]; ok {
		return c.JSON(http.StatusOK, user)
	} else {
		return c.JSON(http.StatusNotFound, UserNotFound)
	}
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	f, err := os.ReadFile(Store)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	s := UserStore{}
	err = json.Unmarshal(f, &s)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	request := UpdateUserRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if _, ok := s.List[id]; !ok {
		return c.JSON(http.StatusNotFound, UserNotFound)
	}

	u := s.List[id]
	u.DisplayName = request.DisplayName
	s.List[id] = u

	b, err := json.Marshal(&s)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = os.WriteFile(Store, b, fs.ModePerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	f, err := os.ReadFile(Store)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	s := UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if _, ok := s.List[id]; !ok {
		return c.JSON(http.StatusNotFound, UserNotFound)
	}

	delete(s.List, id)

	b, err := json.Marshal(&s)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := os.WriteFile(Store, b, fs.ModePerm); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, b)
}
