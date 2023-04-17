package user

import (
	"fmt"
	"refactoring/internal/data"
	"refactoring/internal/services/storage"
	"strconv"
)

type JsonUserCRUD struct {
	fileStorage *storage.FileStorage
}

func NewJsonUserCRUD(fs *storage.FileStorage) *JsonUserCRUD {
	return &JsonUserCRUD{fileStorage: fs}
}

func (s *JsonUserCRUD) GetUsers() (data.UserStore, error) {
	users, err := s.fileStorage.ReadStore()
	return users, err
}

func (s *JsonUserCRUD) GetUser(id string) (data.User, error) {
	users, err := s.fileStorage.ReadStore()
	if err != nil {
		return users.List[id], err
	}
	if user, ok := users.List[id]; ok {
		return user, nil
	} else {
		return data.User{}, fmt.Errorf("Пользователь не найден")
	}
}

func (s *JsonUserCRUD) CreateUser(user data.User) (string, error) {
	users, err := s.fileStorage.ReadStore()
	if err != nil {
		return "", err
	}
	for _, u := range users.List {
		if u.Email == user.Email {
			return "", fmt.Errorf("Пользователь с email '%s' уже существует.", user.Email)
		}
	}
	users.Increment = 1
	for i := 1; i <= len(users.List)+1; i++ {
		stringI := strconv.Itoa(i)
		_, ok := users.List[stringI]
		if !ok {
			users.Increment = i
		}
	}
	user.ID = users.Increment
	users.List[strconv.Itoa(user.ID)] = user

	err = s.fileStorage.WriteStore(users)

	if err != nil {
		return "", err
	}

	return strconv.Itoa(user.ID), nil

}

func (s *JsonUserCRUD) UpdateUser(id string, user data.User) error {

	users, err := s.fileStorage.ReadStore()
	if err != nil {
		return err
	}
	if _, ok := users.List[id]; !ok {
		return fmt.Errorf("пользователь не найден")
	}
	users.List[id] = user
	err = s.fileStorage.WriteStore(users)
	if err != nil {
		return err
	}
	return nil
}

func (s *JsonUserCRUD) DeleteUser(id string) error {
	users, err := s.fileStorage.ReadStore()
	if err != nil {
		return err
	}
	if _, ok := users.List[id]; !ok {
		return fmt.Errorf("пользователь не найден")
	}
	delete(users.List, id)
	err = s.fileStorage.WriteStore(users)
	if err != nil {
		return err
	}
	return nil
}
