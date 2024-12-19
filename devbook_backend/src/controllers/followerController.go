package controller

import (
	"devbook_backend/src/authentication"
	"devbook_backend/src/database"
	"devbook_backend/src/repository"
	"devbook_backend/src/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followedUserId := mux.Vars(r)["userId"]

	followingUserId, err := authentication.GetUserIDFromToken(r)
	if err != nil {
		response.Error(w, "Error getting user id", http.StatusUnauthorized, err.Error())
		return
	}

	if _, err := uuid.Parse(followedUserId); err != nil {
		response.Error(w, "Invalid ID format", http.StatusBadRequest, err.Error())
		return
	}

	if followedUserId == followingUserId {
		response.Error(w, "You can't follow yourself", http.StatusBadRequest, nil)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, "Error trying to connect to the database", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	followersRepository := repository.NewFollowersRepository(db)
	err = followersRepository.Follow(followingUserId, followedUserId)
	if err != nil {
		response.Error(w, "", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followedUserId := mux.Vars(r)["userId"]

	followingUserId, err := authentication.GetUserIDFromToken(r)
	if err != nil {
		response.Error(w, "Error getting user id", http.StatusUnauthorized, err.Error())
		return
	}

	if _, err := uuid.Parse(followedUserId); err != nil {
		response.Error(w, "Invalid ID format", http.StatusBadRequest, err.Error())
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, "Error trying to connect to the database", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	followersRepository := repository.NewFollowersRepository(db)
	err = followersRepository.Unfollow(followingUserId, followedUserId)
	if err != nil {
		response.Error(w, "Unable to unfollow the user", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusNoContent, nil)
}

func GetUserFollowing(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]

	if _, err := uuid.Parse(userId); err != nil {
		response.Error(w, "Invalid ID format", http.StatusBadRequest, err.Error())
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, "Error trying to connect to the database", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	followersRepository := repository.NewFollowersRepository(db)

	following, err := followersRepository.GetUserFollowing(userId)
	if err != nil {
		response.Error(w, "Failed to fetch followers data", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, following)
}

func GetUserFollowers(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]

	if _, err := uuid.Parse(userId); err != nil {
		response.Error(w, "Invalid ID format", http.StatusBadRequest, err.Error())
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, "Error trying to connect to the database", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	followersRepository := repository.NewFollowersRepository(db)

	followers, err := followersRepository.GetUserFollowers(userId)
	if err != nil {
		response.Error(w, "Failed to fetch users who are followed", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, followers)
}
