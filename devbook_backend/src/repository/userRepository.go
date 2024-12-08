package repository

import (
	"database/sql"
	"devbook_backend/src/models"

	"github.com/google/uuid"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (uuid.UUID, error) {
	return uuid.NewUUID()
}
