package controller

import (
	"devbook_backend/src/database"
	"devbook_backend/src/models"
	"devbook_backend/src/repository"
	"devbook_backend/src/response"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, "Error trying to read the request body", http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		response.Error(w, "Error converting request body to JSON", http.StatusBadRequest, err.Error())
		return
	}

	if err = user.Prepare(); err != nil {
		response.Error(w, "Missing required field error", http.StatusUnprocessableEntity, err.Error())
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, "Error trying to connect to the database", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	userRepository := repository.NewUsersRepository(db)

	err = userRepository.Create(user)
	if err != nil {
		response.Error(w, "Error creating user", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, nil)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	if search == "" {
		response.Error(w, "Empty search field", http.StatusNoContent, nil)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, "Error trying to connect to the database", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	userRepository := repository.NewUsersRepository(db)

	result, err := userRepository.Search(search)
	if err != nil {
		response.Error(w, "Error trying search a term", http.StatusInternalServerError, err.Error())
		return
	}

	formatedResult := map[string]interface{}{
		"count":   len(result),
		"results": result,
	}

	response.Success(w, http.StatusOK, formatedResult)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUser"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}
