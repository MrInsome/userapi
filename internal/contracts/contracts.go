package contracts

import (
	"refactoring/internal/data"
)

type FileStorageContract interface {
	ReadStore() (data.UserStore, error)
	WriteStore(store data.UserStore) error
}

type UserStoreContract interface {
	GetUsers() (data.UserStore, error)
	GetUser(id string) (data.User, error)
	CreateUser(user data.User) (string, error)
	UpdateUser(id string, user data.User) error
	DeleteUser(id string) error
}
