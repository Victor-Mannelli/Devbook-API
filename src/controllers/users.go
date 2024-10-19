package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
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
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := user.ParseUserDto(); err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepository := repositories.UsersRepository(db)

	user.ID, err = usersRepository.CreateUser(user)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}
