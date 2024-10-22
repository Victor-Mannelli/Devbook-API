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

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := db.DBConnect()
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepository := repositories.UsersRepository(db)

	savedUser, err := usersRepository.FindUserByEmail(user.Email)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err = utils.ValidateHash(savedUser.Password, user.Password); err != nil {
		utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	jwtToken, _ := utils.CreateTokenJWT(savedUser.ID)

	// fmt.Println(jwtToken)
	// w.Write([]byte(jwtToken))
	utils.HttpJsonResponse(w, http.StatusCreated, jwtToken)
}
