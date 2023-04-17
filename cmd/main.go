package main

import (
	"fmt"
	"net/http"
	"refactoring/internal/api/rest/handlers"
	config2 "refactoring/internal/config"
	"refactoring/internal/services"
	"refactoring/internal/services/storage"
)

func main() {
	config, err := config2.Init("newconfig")
	if err != nil {
		_ = fmt.Errorf("Ошибка инициализации конфига")
	}
	fileStorage := storage.NewFileStorage(config)
	service := services.NewService(fileStorage)
	r := handlers.NewRouter(service, config).Start()
	err = http.ListenAndServe(config.RestPort, r)
	if err != nil {
		_ = fmt.Errorf("Ошибка инициализации http сервера.")
	}
}
