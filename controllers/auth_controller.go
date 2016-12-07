package controllers

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/FEUPTalks/Backend/model"
	"github.com/FEUPTalks/Backend/services"
	"github.com/FEUPTalks/Backend/util"
)

//Login
func Login(writer http.ResponseWriter, request *http.Request) {
	requestUser := new(model.LoginInfo)
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&requestUser)

	token, err := services.Login(requestUser)
	if err != nil {
		log.Println(err)
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(token)
}

//RefreshToken
func RefreshToken(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	requestUser := new(model.User)
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&requestUser)

	writer.Header().Set("Content-Type", "application/json")
	token, err := services.RefreshToken(requestUser)
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	writer.Write(token)
}
