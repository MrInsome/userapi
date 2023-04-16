package user

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"refactoring/internal/config"
	"refactoring/internal/data"
	"strconv"
)

type UserStoreJSON struct {
	*config.Configs
}

func NewUserStoreJSON(c *config.Configs) *UserStoreJSON {
	return &UserStoreJSON{c}
}

func (s *UserStoreJSON) GetUsers() (data.UserStore, error) {
	var users data.UserStore
	StoreFS := http.Dir(s.Directory)
	_, err := StoreFS.Open(s.Name)
	if err == nil {
		f, err := os.ReadFile(s.Name)
		if err != nil {
			return users, err
		}
		err = json.Unmarshal(f, &users)
		if err != nil {
			return users, err
		}
		return users, nil
	} else {
		return data.UserStore{}, err
	}
}

func (s *UserStoreJSON) GetUser(id string) (data.User, error) {
	var users data.UserStore
	StoreFS := http.Dir(s.Directory)
	_, err := StoreFS.Open(s.Name)
	if err == nil {
		f, err := os.ReadFile(s.Name)
		if err != nil {
			return data.User{}, err
		}
		err = json.Unmarshal(f, &users)
		if err != nil {
			return data.User{}, err
		}
		if user, ok := users.List[id]; ok {
			return user, nil
		} else {
			return data.User{}, fmt.Errorf("Пользователь не найден")
		}
	} else {
		return data.User{}, err
	}
}

func (s *UserStoreJSON) CreateUser(user data.User) (string, error) {
	var users data.UserStore
	StoreFS := http.Dir(s.Directory)
	_, err := StoreFS.Open(s.Name)
	if err == nil {
		f, err := os.ReadFile(s.Name)
		if err != nil {
			return "", err
		}
		err = json.Unmarshal(f, &users)
		if err != nil {
			return "", err
		}
		users.Increment = 0
		for _, u := range users.List {
			if u.Email == user.Email {
				return "", fmt.Errorf("Пользователь с email '%s' уже существует.", user.Email)
			}
		}
		for _, u := range users.List {
			if u.ID > users.Increment {
				users.Increment = u.ID - 1
				break
			}
			users.Increment++
		}

		user.ID = users.Increment
		users.List[strconv.Itoa(user.ID)] = user

		b, err := json.Marshal(&users)
		if err != nil {
			return "", err
		}
		err = os.WriteFile(s.Name, b, fs.ModePerm)
		if err != nil {
			return "", err
		}

		return strconv.Itoa(user.ID), nil
	} else {
		return "", err
	}
}

func (s *UserStoreJSON) UpdateUser(id string, user data.User) error {
	StoreFS := http.Dir(s.Directory)
	_, err := StoreFS.Open(s.Name)
	if err == nil {
		f, err := os.ReadFile(s.Name)
		if err != nil {
			return err
		}
		users := data.UserStore{}
		err = json.Unmarshal(f, &users)
		if err != nil {
			return err
		}
		if _, ok := users.List[id]; !ok {
			return fmt.Errorf("пользователь не найден")
		}
		users.List[id] = user
		b, err := json.Marshal(&users)
		if err != nil {
			return err
		}
		err = os.WriteFile(s.Name, b, fs.ModePerm)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func (s *UserStoreJSON) DeleteUser(id string) error {
	StoreFS := http.Dir(s.Directory)
	_, err := StoreFS.Open(s.Name)
	if err == nil {
		f, err := os.ReadFile(s.Name)
		if err != nil {
			return err
		}
		users := data.UserStore{}
		err = json.Unmarshal(f, &users)
		if err != nil {
			return err
		}
		if _, ok := users.List[id]; !ok {
			return fmt.Errorf("пользователь не найден")
		}
		delete(users.List, id)
		b, err := json.Marshal(&users)
		if err != nil {
			return err
		}
		err = os.WriteFile(s.Name, b, fs.ModePerm)
		if err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}
