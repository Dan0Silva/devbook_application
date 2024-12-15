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

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

	if err = user.Prepare("register"); err != nil {
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

func GetUserById(w http.ResponseWriter, r *http.Request) {
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

	userRepository := repository.NewUsersRepository(db)

	result, err := userRepository.SearchById(userId)
	if err != nil {
		response.Error(w, "Error to get user by id", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	var user models.User

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, "Error trying to read the request body", http.StatusBadRequest, err.Error())
		return
	}

	if _, err := uuid.Parse(userId); err != nil {
		response.Error(w, "Invalid ID format", http.StatusBadRequest, err.Error())
		return
	}

	userIdFromToken, err := authentication.GetUserIDFromToken(r)
	if err != nil {
		response.Error(w, "Error getting user id", http.StatusUnauthorized, err.Error())
		return
	}

	if userId != userIdFromToken {
		response.Error(w, "You do not have permission to perform the operation", http.StatusForbidden, nil)
		return
	}

	if err = json.Unmarshal(reqBody, &user); err != nil {
		response.Error(w, "Error converting request body to JSON", http.StatusInternalServerError, err.Error())
		return
	}

	if err = user.Prepare("update"); err != nil {
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
	if err = userRepository.UpdateUser(userId, user); err != nil {
		response.Error(w, "Error trying update user", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]

	if _, err := uuid.Parse(userId); err != nil {
		response.Error(w, "Invalid ID format", http.StatusBadRequest, err.Error())
		return
	}

	userIdFromToken, err := authentication.GetUserIDFromToken(r)
	if err != nil {
		response.Error(w, "Error getting user id", http.StatusUnauthorized, err.Error())
		return
	}

	if userId != userIdFromToken {
		response.Error(w, "You do not have permission to perform the operation", http.StatusForbidden, nil)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, "Error trying to connect to the database", http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	userRepository := repository.NewUsersRepository(db)

	if err = userRepository.Delete(userId); err != nil {
		response.Error(w, "Error to try delete user", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusNoContent, nil)
}
