package repository

import (
	"database/sql"
	"devbook_backend/src/models"
	"fmt"
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

	_, err = statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repository users) Search(search string) ([]models.User, error) {
	var user models.User
	var users []models.User

	search = "%" + search + "%"

	rows, err := repository.db.Query(
		fmt.Sprintf("SELECT ID, NAME, NICK, EMAIL, CREATED_AT FROM USERS WHERE NAME LIKE '%s' OR NICK LIKE '%s';", search, search),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
