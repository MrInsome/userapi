package contracts

import "refactoring/internal/data"

type UserStoreContract interface {
	GetUsers() (map[string]data.User, error)
	GetUser(id string) (data.User, error)
	CreateUser(user data.User) (string, error)
	UpdateUser(id string, user data.User) error
	DeleteUser(id string) error
}
