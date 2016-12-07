package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/model"
	"github.com/FEUPTalks/Backend/util"
	"github.com/gorilla/mux"
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

//Gets all Talk Registrations of the Talk with ID == talkID
func (*TalkRegistrationController) GetTalkRegistrationsWithTalkID(writer http.ResponseWriter, request *http.Request) {
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(request)
	talkID, err := strconv.Atoi(vars["talkID"])
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	talkRegistrations, err := instance.GetTalkRegistrationsWithTalkID(talkID)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendJSON(
		writer,
		request,
		talkRegistrations,
		http.StatusOK,
	)
}
