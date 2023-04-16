package main

import (
	"errors"
	"net/http"
	"refactoring/internal/api/rest/handlers"
)

var (
	UserNotFound = errors.New("user_not_found")
)

func main() {
	r := handlers.NewRouter()
	http.ListenAndServe(":3333", r)
}
