package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/model"
)

//TalkController struct
type TalkRegistrationController struct {
}

//Creates a new Talk Registration
func (*TalkRegistrationController) Create(writer http.ResponseWriter, request *http.Request) {
	talkRegistrationToCreate := model.NewTalkRegistration()
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&talkRegistrationToCreate)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	instance.SaveTalkRegistration(talkRegistrationToCreate)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusCreated)
}
