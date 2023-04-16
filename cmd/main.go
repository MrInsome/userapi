package main

import (
	"fmt"
	"net/http"
	"refactoring/internal/api/rest/handlers"
	config2 "refactoring/internal/config"
	"refactoring/internal/services"
)

func main() {
	config, err := config2.Init("newconfig")
	if err != nil {
		fmt.Errorf("Ошибка инициализации конфига")
	}
	service := services.NewService(config)
	r := handlers.NewRouter(service, config).Start()
	http.ListenAndServe(config.RestPort, r)
}
