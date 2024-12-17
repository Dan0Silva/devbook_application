package repository

import (
	"database/sql"
	"devbook_backend/src/models"
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
	statement, err := repository.db.Prepare("delete from FOLLOWERS where FOLLOWING_ID = ? and FOLLOWED_ID = ?;")
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

// usuários que o userId está seguindo
func (repository followers) GetUserFollowing(userId string) ([]models.User, error) {
	var user models.User
	var usersList []models.User

	rows, err := repository.db.Query(`
		SELECT ID, NAME, NICK, EMAIL FROM USERS JOIN FOLLOWERS F ON F.FOLLOWED_ID = USERS.ID WHERE F.FOLLOWING_ID = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email); err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}

	return usersList, nil
}

// usuários que estao seguindo o userId
func (repository followers) GetUserFollowers(userId string) ([]models.User, error) {
	var user models.User
	var usersList []models.User

	rows, err := repository.db.Query(`
		SELECT ID, NAME, NICK, EMAIL FROM USERS JOIN FOLLOWERS F ON F.FOLLOWING_ID = USERS.ID WHERE F.FOLLOWED_ID = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email); err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}

	return usersList, nil
}
