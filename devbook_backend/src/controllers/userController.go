package controller

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUser"))
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
