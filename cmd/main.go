package main

import (
	"fmt"
	"net/http"
	"refactoring/internal/api/rest/handlers"
	config2 "refactoring/internal/config"
)

func main() {
	config, err := config2.Init("newconfig")
	if err != nil {
		fmt.Errorf("Не смог найти конфиг")
	}
	r := handlers.NewRouter()
	http.ListenAndServe(config.RestPort, r)
}
