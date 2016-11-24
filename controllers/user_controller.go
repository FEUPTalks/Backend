package controllers

import (
	"encoding/json"

	"log"
	"net/http"

	"strconv"

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

	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	instance.SaveUser(userToCreate)

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
