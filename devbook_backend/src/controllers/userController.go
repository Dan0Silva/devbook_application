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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllUsers"))
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
