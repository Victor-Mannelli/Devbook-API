package controllers

import (
	"api/src/db"
	"api/src/modules"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FindAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindAllUser"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindUser"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	utils.CheckError(err)

	var user modules.User
	err = json.Unmarshal(requestBody, &user)
	utils.CheckError(err)

	db, err := db.DBConnect()
	utils.CheckError(err)

	usersRepository := repositories.UsersRepository(db)

	userId, err := usersRepository.CreateUser(user)
	utils.CheckError(err)

	w.Write([]byte(fmt.Sprintf("Created user with id: %d", userId)))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}
