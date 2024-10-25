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

	"github.com/gorilla/mux"
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

	post.AuthorId = userIdFromToken

	if err = post.ParsePostDto("createPost"); err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postsRepository := repositories.PostsRepository(db)

	post.PostId, err = postsRepository.CreatePost(post)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusCreated, post)
}

func FindPosts(w http.ResponseWriter, r *http.Request) {
	userIdFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postsRepository := repositories.PostsRepository(db)

	posts, err := postsRepository.FindPosts(userIdFromToken)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusOK, posts)
}

func FindPostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//* turning userid string param value to int
	userId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
	}

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postsRepository := repositories.PostsRepository(db)

	user, err := postsRepository.FindPostById(userId)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusCreated, user)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userIdFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["postId"], 10, 64)
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

	postsRepository := repositories.PostsRepository(db)
	savePost, err := postsRepository.FindPostById(postId)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if savePost.AuthorId != userIdFromToken {
		utils.HttpErrorResponse(w, http.StatusForbidden, errors.New("access forbidden for current action"))
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

	if err = post.ParsePostDto("updatePost"); err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err = postsRepository.UpdatePost(postId, post); err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusCreated, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	userIdFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["postId"], 10, 64)
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

	postsRepository := repositories.PostsRepository(db)

	savedPost, err := postsRepository.FindPostById(postId)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if savedPost.AuthorId != userIdFromToken {
		utils.HttpErrorResponse(w, http.StatusForbidden, errors.New("access forbidden for current action"))
		return
	}

	if err = postsRepository.DeletePost(postId); err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(w, http.StatusNoContent, nil)
}
