package services

import "github.com/MraGLO/practica/internal/repository"

type Services struct {
	Database *DatabaseService
}

func MakeServices(repo *repository.Repository) *Services {
	return &Services{Database: newDatabaseService(repo.Database)}
}
