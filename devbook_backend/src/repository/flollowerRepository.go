package repository

import (
	"database/sql"
	"errors"
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

func (repository followers) Unfollow(followingUserId, followedUserId string) error {
	statement, err := repository.db.Prepare("delete ignore from FOLLOWERS where FOLLOWING_ID = ? and FOLLOWED_ID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(followingUserId, followedUserId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were deleted, check if the relationship exists")
	}

	return nil
}
