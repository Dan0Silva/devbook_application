package repository

import (
	"database/sql"
	"devbook_backend/src/models"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) error {
	statement, err := repository.db.Prepare("insert into USERS (NAME, NICK, EMAIL, PASSWORD) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	t, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return err
	}

	t.LastInsertId()

	return nil
}
