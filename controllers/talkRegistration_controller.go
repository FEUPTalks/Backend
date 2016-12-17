package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"database/sql"

	"github.com/FEUPTalks/Backend/core/authentication"
	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/model"
	"github.com/FEUPTalks/Backend/services"
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

	exists, err := checkIfRegistrationExists(talkRegistrationToCreate)
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {

		err = instance.CreateTemporaryTalkRegistrationChange(talkRegistrationToCreate)
		if err != nil {
			util.ErrHandler(err, writer, http.StatusInternalServerError)
			return
		}

		err = sendEmailForTalkRegistrationReplacementConfirmation(talkRegistrationToCreate)
		writer.WriteHeader(http.StatusOK)
	} else {
		talkRegistrationLogToCreate := model.NewTalkRegistrationLogWithTalkRegistration(talkRegistrationToCreate, 0)

		instance.SaveTalkRegistration(talkRegistrationToCreate)
		instance.SaveTalkRegistrationLog(talkRegistrationLogToCreate)
		writer.WriteHeader(http.StatusCreated)
	}

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

//ConfirmTalkRegistrationChange used to confirm talk registration change
func (*TalkRegistrationController) ConfirmTalkRegistrationChange(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	talkID, err := strconv.Atoi(vars["talkID"])
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	token := request.URL.Query().Get("token")
	if token == "" {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	authBackend, err := authentication.GetJWTAuthenticationBackend()
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	email, err := authBackend.GetTokenClaim(token, "sub")
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	talkRegistrations, err := instance.GetTemporaryTalkRegistrationsWithTalkID(talkID)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var talkRegistration *model.TalkRegistration

	for _, tr := range talkRegistrations {
		if tr.Email == email {
			talkRegistration = tr
			break
		}
	}

	if talkRegistration == nil {
		util.ErrHandler(errors.New("Talk registration does not exist for this user in this talk"), writer, http.StatusNotFound)
		return
	}

	instance.EditTalkRegistration(talkID, email, talkRegistration)

	writer.WriteHeader(http.StatusOK)
}

func checkIfRegistrationExists(talkRegistrationToCreate *model.TalkRegistration) (bool, error) {
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		return false, err
	}

	check, err := instance.CheckIfRegistrationExistsInTalk(talkRegistrationToCreate.Email, talkRegistrationToCreate.TalkID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return check, nil
}

func sendEmailForTalkRegistrationReplacementConfirmation(replacementOfTalkRegistration *model.TalkRegistration) error {

	authBackend, err := authentication.GetJWTAuthenticationBackend()
	if err != nil {
		return err
	}

	token, err := authBackend.GenerateToken(replacementOfTalkRegistration.Email)
	if err != nil {
		return err
	}

	confirmationLink := url.URL{}

	confirmationLink.Scheme = "http"
	confirmationLink.Host = "les16b.fe.up.pt:8144"
	confirmationLink.Path = "/talkRegistration/" + strconv.Itoa(replacementOfTalkRegistration.TalkID) + "/confirm"
	confirmationLink.RawQuery = "token=" + token

	emailToSend := &model.Email{}

	emailToSend.EmailTo = replacementOfTalkRegistration.Email
	emailToSend.NameTo = replacementOfTalkRegistration.Name
	emailToSend.URL = confirmationLink.String()

	err = services.SendEmailConfirmation(emailToSend, services.Link)
	if err != nil {
		return err
	}

	return nil
}
