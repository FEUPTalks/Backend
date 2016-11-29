package controllers

import (
	"log"
	"net/http"

	"encoding/json"

	"strconv"

	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/model"
	"github.com/FEUPTalks/Backend/util"
	"github.com/gorilla/mux"
)

//TalkController struct
type TalkController struct {
}

//Index func return all talks in the database
func (*TalkController) Index(writer http.ResponseWriter, request *http.Request) {
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	talks, err := instance.GetAllTalks()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendJSON(
		writer,
		request,
		talks,
		http.StatusOK,
	)
}

//Create creates a new Talk
func (*TalkController) Create(writer http.ResponseWriter, request *http.Request) {
	talkToCreate := model.NewTalk()
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&talkToCreate)
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
	instance.SaveTalk(talkToCreate)

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
}

//GetTalk func return talk with specific id from the database
func (*TalkController) GetTalk(writer http.ResponseWriter, request *http.Request) {
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(request)
	talkID, err := strconv.Atoi(vars["talkID"])
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}
	talk, err := instance.GetTalk(talkID)
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	util.SendJSON(
		writer,
		request,
		talk,
		http.StatusOK,
	)
}

//SetTalk func update database to a specific talkID
func (*TalkController) SetTalk(writer http.ResponseWriter, request *http.Request) {
	talkToCreate := model.NewTalk()
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&talkToCreate)
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
	instance.SetTalk(talkToCreate)

	writer.WriteHeader(http.StatusOK)
}
