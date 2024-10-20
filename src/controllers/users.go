package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("filter"))

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.UsersRepository(db)

	users, err := userRepository.FindFilteredUsers(nameOrUsername)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusOK, users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// turning userid string param value to int
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.UsersRepository(db)

	user, err := userRepository.FindUser(userId)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusCreated, user)
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

	if err := user.ParseUserDto("createUser"); err != nil {
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
	params := mux.Vars(r)

	// turning userid string param value to int
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var updatedUserDto models.User
	if err := json.Unmarshal(requestBody, &updatedUserDto); err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = updatedUserDto.ParseUserDto("updateUser"); err != nil {
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

	err = usersRepository.UpdateUser(userId, updatedUserDto)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}
