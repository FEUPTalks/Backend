package controllers

import (
	"encoding/json"

	"log"
	"net/http"
	"net/smtp"

	"fmt"
	"strconv"

	"strings"

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
	subject := "FEUPTalks Validation"
	from := "feuptalks@gmail.com"
	//body := ParseTemplateHTML(user.Name, fmt.Sprintf("http://%s%s/%s", request.Host, request.URL.Path, user.HashCode))
	body := ParseTemplate(user.Name, fmt.Sprintf("http://%s%s/%s", request.Host, request.URL.Path, user.HashCode))
	msg := "Subject: " + subject + "\n" +
		"From: " + from + "\n" +
		"To: " + user.Email + "\n" +
		body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"feuptalks@gmail.com",
		[]string{user.Email},
		[]byte(msg),
	)
	return err
}

//ParseTemplateHTML fill template html
func ParseTemplateHTML(name string, urlValid string) string {
	template :=
		(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN""http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html></head><body><p>Hello STRNAME
	<a href="STRURL">Validation address</a>
	</p></body></html>`)
	template = strings.Replace(template, "STRNAME", name, -1)
	template = strings.Replace(template, "STRURL", urlValid, -1)
	return template
}

//ParseTemplate fill template in Text
func ParseTemplate(name string, urlValid string) string {
	template := (`
Hello STRNAME, 
		
To access FEUPtalks you don't need any passoword. You just need to follow this link:
"STRURL"

FEUPTalks`)
	template = strings.Replace(template, "STRNAME", name, -1)
	template = strings.Replace(template, "STRURL", urlValid, -1)
	return template
}
