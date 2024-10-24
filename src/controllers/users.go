package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
	"errors"
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

func FindUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//* turning userid string param value to int
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

	user, err := userRepository.FindUserById(userId)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusCreated, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
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

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	userIdFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var updatePasswordDto models.Password
	if err = json.Unmarshal(reqBody, &updatePasswordDto); err != nil {
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

	savedHashedPassword, err := usersRepository.FindUsersPassword(userIdFromToken)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err = utils.ValidateHash(savedHashedPassword, updatePasswordDto.Password); err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, errors.New("passwords don't match"))
	}

	hashedPassword, err := utils.HashString(updatePasswordDto.NewPassword)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = usersRepository.UpdatePassword(userIdFromToken, string(hashedPassword)); err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// utils.HttpJsonResponse(w, http.StatusOK, nil)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//* turning userid string param value to int
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}
	//* checks if user is the owner of what's being changed
	if userIdFromToken != userId {
		utils.HttpErrorResponse(w, http.StatusForbidden, errors.New("access forbidden for current action"))
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var updatedUserDto models.User
	if err := json.Unmarshal(reqBody, &updatedUserDto); err != nil {
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
	params := mux.Vars(r)

	//* turning userid string param value to int
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}
	//* checks if user is the owner of what's being changed
	if userIdFromToken != userId {
		utils.HttpErrorResponse(w, http.StatusForbidden, errors.New("access forbidden for current action"))
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepository := repositories.UsersRepository(db)
	if err = usersRepository.DeleteUser(userId); err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusNoContent, nil)
}
