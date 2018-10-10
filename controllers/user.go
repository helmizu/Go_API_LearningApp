package controllers

import (
	"encoding/json"
	"net/http"

	"PW/config"
	"PW/library"
	"PW/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var conf = config.Config{}
var lib = library.UserCtr{}
var res = library.Utill{}

// SemuaGuru use to GET all users
func SemuaGuru(w http.ResponseWriter, r *http.Request) {
	guru, err := lib.FindAllGuru()
	if err != nil {
		res.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if guru == nil {
		guru = []models.GuruPrimary{}
	}
	res.RespondWithJSON(w, http.StatusOK, guru)
}

// CariGuru use to GET a users by its ID
func CariGuru(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	guru, err := lib.FindGuruByID(params["id"])
	if err != nil {
		res.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	res.RespondWithJSON(w, http.StatusOK, guru)
}

// DaftarGuru use to POST a new user
func DaftarGuru(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var guru models.GuruPrimary
	if err := json.NewDecoder(r.Body).Decode(&guru); err != nil {
		res.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	guru.ID = bson.NewObjectId()
	if err := lib.SignupGuru(guru); err != nil {
		res.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	res.RespondWithJSON(w, http.StatusCreated, guru)
}

// UpdateGuru use to PUT update an existing user
func UpdateGuru(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var guru models.GuruPrimary

	guru.ID = bson.ObjectIdHex(params["id"])
	if err := json.NewDecoder(r.Body).Decode(&guru); err != nil {
		res.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := lib.UpdateGuru(guru); err != nil {
		res.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	res.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteGuru use to DELETE an existing user
func DeleteGuru(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	guru, err := lib.FindGuruByID(params["id"])
	if err != nil {
		res.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	if err := lib.DeleteGuru(guru); err != nil {
		res.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	res.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
