package controllers

import (
	"log"
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/RAyres23/back-end/model"
	"github.com/RAyres23/back-end/util"
)

// TalkController struct
type TalkController struct {
}

// Index func return all talks in database
func (c *TalkController) Index(writer http.ResponseWriter, request *http.Request) {
	util.SendJSON(
		writer,
		request,
		[]*model.Talk{model.NewEmptyTalk()},
		http.StatusOK,
	)
}

// Create creates a new Talk
func (c *TalkController) Create(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	talkToCreate := model.NewEmptyTalk()
	decoder := json.NewDecoder(request.Body)
	error := decoder.Decode(&talkToCreate)
	if error != nil {
		log.Println(error.Error())
		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}
	outgoingJSON, err := json.Marshal(talkToCreate)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusCreated)
	fmt.Fprint(writer, string(outgoingJSON))
}
