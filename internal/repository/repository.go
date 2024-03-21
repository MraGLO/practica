package repository

import (
	//"github.com/MraGLO/practica/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
}

type Repository struct {
	Database Database
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{Database: newDatabaseRepo(db)}
}
