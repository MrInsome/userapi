package data

import "time"

type User struct {
	ID          int
	CreatedAt   time.Time `json:"created_at"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
}

type UserList map[string]User

type UserStore struct {
	Increment int      `json:"increment"`
	List      UserList `json:"list"`
}

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}
