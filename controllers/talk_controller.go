package controllers

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/RAyres23/LESTeamB-backend/database"
	"github.com/RAyres23/LESTeamB-backend/model"
	"github.com/RAyres23/LESTeamB-backend/util"
)

// TalkController struct
type TalkController struct {
}

// Index func return all talks in database
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
	writer.Header().Set("Content-Type", "application/json")

	talkToCreate := &model.Talk{}
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
}
