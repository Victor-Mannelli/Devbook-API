package controllers

import "net/http"

func FindAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindAllUser"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindUser"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUser"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}
