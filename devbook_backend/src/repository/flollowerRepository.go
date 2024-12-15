package repository

import (
	"database/sql"
)

type followers struct {
	db *sql.DB
}

func NewFollowersRepository(db *sql.DB) *followers {
	return &followers{db}
}

func (repository followers) Follow(followingUserId, followedUserId string) error {
	statement, err := repository.db.Prepare("insert ignore into FOLLOWERS (FOLLOWING_ID, FOLLOWED_ID) values (?, ?);")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(followingUserId, followedUserId); err != nil {
		return err
	}

	return nil
}
