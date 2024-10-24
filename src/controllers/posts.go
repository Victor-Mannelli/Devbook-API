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

func CreatePost(w http.ResponseWriter, r *http.Request) {
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

	var post models.Post
	if err := json.Unmarshal(reqBody, &post); err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postsRepository := repositories.UsersRepository(db)

	// savedHashedPassword, err := usersRepository.FindUsersPassword(userIdFromToken)
	// if err != nil {
	// 	utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
	// 	return
	// }

}
func FindPosts(w http.ResponseWriter, r *http.Request) {

}
func FindPostById(w http.ResponseWriter, r *http.Request) {

}
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
