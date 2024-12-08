package controller

import (
	"devbook_backend/src/database"
	"devbook_backend/src/models"
	"devbook_backend/src/repository"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao tentar ler o corpo da requisição"))
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Erro ao converter corpo da requisição para json"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao tentar coneção com o banco"))
		return
	}
	defer db.Close()

	userRepository := repository.NewUsersRepository(db)

	_, err = userRepository.Create(user)
	if err != nil {
		w.Write([]byte("Erro ao criar usuário"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário criado com sucesso"))

	// statement, err := db.Prepare("insert into USERS ")
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
