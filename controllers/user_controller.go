package controllers

import (
	"encoding/json"

	"log"
	"net/http"
	"net/smtp"

	"strconv"

	"fmt"

	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/model"
	"github.com/FEUPTalks/Backend/util"
)

//UserController struct
type UserController struct {
}

//Create creates a new User
func (*UserController) Create(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	userToCreate := model.NewUser()
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userToCreate)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	userToCreate.SetNewHashCode()

	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	instance.SaveUser(userToCreate)

	err = SendEmailValidation(userToCreate, request)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

//GetUser func return user with specific hashCodeValidation
func (*UserController) GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}
	ctx := request.Context()

	hashcode := ctx.Value("hashcode")
	userIDstr := ctx.Value("userID")

	user := model.NewUser()

	if hashcode != nil {
		user, err = instance.GetUserByhash(hashcode.(string))
	} else {
		userID, err := strconv.Atoi(userIDstr.(string))
		if err != nil {
			util.ErrHandler(err, writer, http.StatusInternalServerError)
			return
		}
		user, err = instance.GetUserByID(userID)
	}
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	util.SendJSON(
		writer,
		request,
		user,
		http.StatusOK,
	)
}

//SetUser func update database to a specific talkID
func (*UserController) SetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	userToCreate := model.NewUser()
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userToCreate)
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
	instance.SetUser(userToCreate)

	writer.WriteHeader(http.StatusOK)
}

//DeleteLastUser delete user created in tests
func (*UserController) DeleteLastUser() {
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		return
	}
	instance.DeleteLastUser()
}

// SendEmailValidation sent an email to user in order to get the validation link
func SendEmailValidation(user *model.User, request *http.Request) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"feuptalks@gmail.com",
		"Talks123",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"feuptalks@gmail.com",
		[]string{user.Email},
		[]byte(fmt.Sprintf("http://%s%s/user/%s\n", request.Host, request.URL.Path, user.HashCode)),
	)
	return err
}
