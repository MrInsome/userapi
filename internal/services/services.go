package services

import (
	"refactoring/internal/contracts"
	"refactoring/internal/services/storage"
	"refactoring/internal/services/user"
)

type Service struct {
	contracts.FileStorageContract
	contracts.UserStoreContract
}

func NewService(fileStorage *storage.FileStorage) *Service {
	return &Service{
		FileStorageContract: storage.NewFileStorage(fileStorage.Configs),
		UserStoreContract:   user.NewJsonUserCRUD(fileStorage),
	}
}
