package services

import (
	"refactoring/internal/config"
	"refactoring/internal/contracts"
	"refactoring/internal/services/user"
)

type Service struct {
	contracts.UserStoreContract
}

func NewService(conf *config.Configs) *Service {
	return &Service{UserStoreContract: user.NewUserStoreJSON(conf)}
}
