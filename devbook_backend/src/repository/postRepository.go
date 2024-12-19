package repository

import (
	"database/sql"
	"devbook_backend/src/models"

	"github.com/google/uuid"
)

type posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (repository posts) Create(userId string, post models.Post) (*models.Post, error) {
	rows, err := repository.db.Query("select NICK from USERS where ID = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nick string

	for rows.Next() {
		if err = rows.Scan(&nick); err != nil {
			return nil, err
		}
	}

	statement, err := repository.db.Prepare("insert into POSTS (TITLE, CONTENT, AUTHOR_ID, AUTHOR_NICK) values (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	_, err = statement.Exec(post.Title, post.Content, userId, nick)
	if err != nil {
		return nil, err
	}

	post.AuthorId, err = uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	post.AuthorNick = nick

	return &post, nil
}
