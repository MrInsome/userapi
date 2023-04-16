package services

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"refactoring/internal/data"
	"strconv"
)

type UserStoreJSON struct {
	storePath string
}

func NewUserStoreJSON(storePath string) *UserStoreJSON {
	return &UserStoreJSON{storePath}
}

func (s *UserStoreJSON) GetUsers() (map[string]data.User, error) {
	f, err := os.ReadFile(s.storePath)
	if err != nil {
		return nil, err
	}
	users := data.UserStore{}
	err = json.Unmarshal(f, &users)
	if err != nil {
		return nil, err
	}
	return users.List, nil
}

func (s *UserStoreJSON) GetUser(id string) (data.User, error) {
	f, err := os.ReadFile(s.storePath)
	if err != nil {
		return data.User{}, err
	}
	users := data.UserStore{}
	err = json.Unmarshal(f, &users)
	if err != nil {
		return data.User{}, err
	}
	if user, ok := users.List[id]; ok {
		return user, nil
	} else {
		return data.User{}, fmt.Errorf("пользователь не найден")
	}
}

func (s *UserStoreJSON) CreateUser(user data.User) (string, error) {
	f, err := os.ReadFile(s.storePath)
	if err != nil {
		return "", err
	}
	users := data.UserStore{}
	err = json.Unmarshal(f, &users)
	if err != nil {
		return "", err
	}

	for _, u := range users.List {
		if u.Email == user.Email {
			return "", fmt.Errorf("Пользователь с email '%s' уже существует.", user.Email)
		}
	}

	users.Increment++
	user.ID = users.Increment
	users.List[strconv.Itoa(user.ID)] = user

	b, err := json.Marshal(&users)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(s.storePath, b, fs.ModePerm)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(user.ID), nil
}

func (s *UserStoreJSON) UpdateUser(id string, user data.User) error {
	f, err := os.ReadFile(s.storePath)
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

	if user.Email != "" {
		for _, u := range users.List {
			if u.ID != user.ID && u.Email == user.Email {
				return fmt.Errorf("Пользователь с email '%s' уже существует.", u.Email)
			}
		}
	}
	users.List[id] = user
	b, err := json.Marshal(&users)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.storePath, b, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStoreJSON) DeleteUser(id string) error {
	f, err := os.ReadFile(s.storePath)
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
	err = os.WriteFile(s.storePath, b, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
