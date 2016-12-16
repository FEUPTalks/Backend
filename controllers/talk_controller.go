package controllers

import (
	"log"
	"net/http"

	"encoding/json"

	"strconv"

	"errors"

	"github.com/FEUPTalks/Backend/core/authentication"
	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/model"
	"github.com/FEUPTalks/Backend/model/talkState"
	"github.com/FEUPTalks/Backend/settings"
	//"github.com/FEUPTalks/Backend/model/talkState/talkStateFactory"
	"github.com/FEUPTalks/Backend/model/talkState/talkStateFactory"
	"github.com/FEUPTalks/Backend/util"
	"github.com/gorilla/mux"
)

//TalkController struct
type TalkController struct {
}

//All func return all talks in the database
func (*TalkController) All(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {

	authBackend, err := authentication.GetJWTAuthenticationBackend()
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	email, err := authBackend.ExtractEmail(request)
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	if !settings.IsAdminOrEmployee(email) {
		util.ErrHandler(errors.New("Not allowed"), writer, http.StatusUnauthorized)
		return
	}

	var talks []*model.Talk

	state := request.FormValue("state")
	if state != "" {
		talks, err = getTalksWithState(state)
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		instance, err := database.GetTalkDatabaseManagerInstance()
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		talks, err = instance.GetAllTalks()
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	util.SendJSON(
		writer,
		request,
		talks,
		http.StatusOK,
	)
}

//Index func return all published talks in the database
func (*TalkController) Index(writer http.ResponseWriter, request *http.Request) {
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	talks, err := instance.GetTalksWithState(&talkState.PublishedTalkState{})
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

	/* This only happens if :
		- the role is 3 (public user)
		- the role is 2 (employee) and state is not accepted (waiting room)

	if talk.StateValue < talkStateFactory.GetPublishedTalkStateValue() {
		authBackend, err := authentication.GetJWTAuthenticationBackend()
		if err != nil {
			util.ErrHandler(err, writer, http.StatusInternalServerError)
			return
		}

		email, err := authBackend.ExtractEmail(request)
		if err != nil {
			util.ErrHandler(err, writer, http.StatusInternalServerError)
			return
		}

		if !settings.IsAdminOrEmployee(email) {
			util.ErrHandler(errors.New("Not allowed"), writer, http.StatusUnauthorized)
			return
		}
	}
	*/

	util.SendJSON(
		writer,
		request,
		talk,
		http.StatusOK,
	)
}

//SetTalk func update database to a specific talkID
func (*TalkController) SetTalk(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	talkToCreate := model.NewTalk()
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&talkToCreate)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(request)
	talkToCreate.TalkID, err = strconv.Atoi(vars["talkID"])
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	authBackend, err := authentication.GetJWTAuthenticationBackend()
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	email, err := authBackend.ExtractEmail(request)
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	if !settings.IsAdmin(email) {
		talk, err := instance.GetTalk(talkToCreate.TalkID)
		if err != nil {
			util.ErrHandler(err, writer, http.StatusUnauthorized)
			return
		}
		if email != talk.ProponentEmail {
			util.ErrHandler(errors.New("Not authorized"), writer, http.StatusUnauthorized)
			return
		}
	}

	instance.SetTalk(talkToCreate)

	writer.WriteHeader(http.StatusOK)
}

//SetTalkState update database talk input talkid and state to change
func (*TalkController) SetTalkState(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	log.Println("SetTalkState")
	vars := mux.Vars(request)
	talkID, err := strconv.Atoi(vars["talkID"])
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}
	newState, err := strconv.Atoi(request.URL.Query().Get("state"))
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}
	if newState == 0 {
		http.Error(writer, "State=0", http.StatusInternalServerError)
		return
	}
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	instance.SetTalkState(talkID, newState)
	writer.WriteHeader(http.StatusOK)
}

//SetTalkRoom update database talk input talkid and room to change
func (*TalkController) SetTalkRoom(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	log.Println("SetTalkRoom")
	vars := mux.Vars(request)
	talkID, err := strconv.Atoi(vars["talkID"])
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}
	room := request.URL.Query().Get("room")
	if room == "" {
		http.Error(writer, "Room=null", http.StatusInternalServerError)
		return
	}
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	instance.SetTalkRoom(talkID, room)
	writer.WriteHeader(http.StatusOK)
}

func getTalksWithState(state string) ([]*model.Talk, error) {
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var talks []*model.Talk

	i, err := strconv.ParseInt(state, 10, 8)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Invalid state")
	}

	stateObj, err := talkStateFactory.GetTalkState(uint8(i))
	if err != nil {
		log.Println(err)
		return nil, errors.New("Invalid state")
	}

	talks, err = instance.GetTalksWithState(stateObj)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return talks, nil
}

//DeleteLastTalk delete user created in tests
func (*TalkController) DeleteLastTalk() {
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		return
	}
	instance.DeleteLastTalk()
}

//GetLastTalkID delete user created in tests
func (*TalkController) GetLastTalkID() int {
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		return -1
	}
	id, err := instance.GetLastTalkID()
	if err != nil {
		log.Println(err)
		return -1
	}
	return id
}