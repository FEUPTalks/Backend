package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/util"
	"github.com/gorilla/mux"
)

//TalkController struct
type TalkRegistrationLogController struct {
}

//Gets all Talk Registrations of the Talk with ID == talkID
func (*TalkRegistrationLogController) GetTalkRegistrationLogsWithTalkID(writer http.ResponseWriter, request *http.Request) {
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

	talkRegistrationLogs, err := instance.GetTalkRegistrationLogsWithTalkID(talkID)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendJSON(
		writer,
		request,
		talkRegistrationLogs,
		http.StatusOK,
	)
}
