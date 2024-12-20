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

func (repository users) Delete(id string) error {
	statement, err := repository.db.Prepare("DELETE FROM USERS WHERE ID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository users) GetByEmail(email string) (*models.User, error) {
	var user models.User

	rows, err := repository.db.Query("SELECT * FROM USERS WHERE EMAIL = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt); err != nil {
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

func (repository users) GetUserPassword(id string) (*string, error) {
	var password string

	rows, err := repository.db.Query("SELECT PASSWORD FROM USERS WHERE ID = ?;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&password); err != nil {
			return nil, err
		}
	}

	return &password, err
}

func (repository users) UpdatePassoword(id string, password string) error {
	statement, err := repository.db.Prepare(`UPDATE USERS SET PASSWORD = ? WHERE ID = ?;`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(password, id)
	if err != nil {
		return err
	}

	return nil
}
