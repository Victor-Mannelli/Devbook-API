package controllers

import (
	"api/src/db"
	"api/src/modules"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	if err != nil {
		log.Fatal(err)
	}

	var user modules.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	usersRepository := repositories.UsersRepository(db)

	userId, err := usersRepository.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Created user with id: %d", userId)))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}
