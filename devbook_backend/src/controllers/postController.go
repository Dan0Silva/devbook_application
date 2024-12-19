package controller

import (
	"devbook_backend/src/authentication"
	"devbook_backend/src/database"
	"devbook_backend/src/models"
	"devbook_backend/src/repository"
	"devbook_backend/src/response"
	"encoding/json"
	"io"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	userId, err := authentication.GetUserIDFromToken(r)
	if err != nil {
		response.Error(w, "Error getting user id", http.StatusUnauthorized, err.Error())
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, "Error trying to read the request body", http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := json.Unmarshal(reqBody, &post); err != nil {
		response.Error(w, "Error converting request body to JSON", http.StatusBadRequest, err.Error())
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, "Error trying to connect to the database", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	followerRepository := repository.NewPostRepository(db)

	createdPost, err := followerRepository.Create(userId, post)
	if err != nil {
		response.Error(w, "Error to create post", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusNoContent, createdPost)
}
