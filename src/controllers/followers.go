package controllers

import (
	"api/src/db"
	"api/src/repositories"
	"api/src/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//* turning userid string param value to int
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	followersRepository := repositories.UsersRepository(db)

	followers, err := followersRepository.FindFollowers(userId)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusCreated, followers)
}

func FindFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//* turning userid string param value to int
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	followersRepository := repositories.UsersRepository(db)

	following, err := followersRepository.FindFollowing(userId)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusCreated, following)
}

func Follow(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.UserIdFromToken(r)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userToFollow, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if userToFollow == userId {
		utils.HttpErrorResponse(w, http.StatusForbidden, errors.New("you can not follow yourself"))
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	followersRepository := repositories.FollowersRepository(db)

	if err = followersRepository.Follow(userId, userToFollow); err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusOK, nil)
}

func UnFollow(w http.ResponseWriter, r *http.Request) {
	followerId, err := utils.UserIdFromToken(r)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		utils.HttpErrorResponse(w, http.StatusForbidden, errors.New("you can unfollow yourself"))
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	followersRepository := repositories.FollowersRepository(db)

	if err = followersRepository.UnFollow(followerId, userId); err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusOK, nil)
}
