package repository

import (
	"database/sql"
	"devbook_backend/src/models"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
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

func (repository users) SearchById(id string) (*models.User, error) {
	var user models.User

	rows, err := repository.db.Query("SELECT ID, NAME, NICK, EMAIL, CREATED_AT FROM USERS WHERE ID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
	}

	defaultTime := time.Time{}
	defaultUUID := uuid.UUID{}

	if user.Id == defaultUUID && user.CreatedAt == defaultTime {
		return nil, errors.New("could not find any user")
	}

	return &user, nil
}

func (repository users) UpdateUser(id string, user models.User) error {
	statement, err := repository.db.Prepare("UPDATE USERS SET NAME = ?, NICK = ?, EMAIL = ? WHERE ID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nick, user.Email, id)
	if err != nil {
		return err
	}

	return nil
}
