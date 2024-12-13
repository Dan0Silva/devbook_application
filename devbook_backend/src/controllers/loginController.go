package controller

import (
	"devbook_backend/src/authentication"
	"devbook_backend/src/database"
	"devbook_backend/src/models"
	"devbook_backend/src/repository"
	"devbook_backend/src/response"
	"devbook_backend/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, "Error trying to read the request body", http.StatusUnprocessableEntity, err.Error())
		return
	}

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

	userFromDB, err := userRepository.GetByEmail(user.Email)
	if err != nil {
		response.Error(w, "Error to try sign in", http.StatusInternalServerError, err.Error())
		return
	}

	if err = security.VerifyPassword(user.Password, userFromDB.Password); err != nil {
		response.Error(w, "Password incorrect", http.StatusUnauthorized, err.Error())
		return
	}

	userToken, err := authentication.CreateToken(user.Id)
	if err != nil {
		response.Error(w, "Error to create user token", http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, userToken)
}
