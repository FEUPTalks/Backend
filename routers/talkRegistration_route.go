package routers

import (
	"github.com/FEUPTalks/Backend/controllers"
	"github.com/gorilla/mux"
)

const (
	talkRegistrationPath string = "/talkRegistration"
)

//Sets the routes for /talkRegistration functionalities
func SetTalkRegistrationRoutes(router *mux.Router) *mux.Router {

	talkRegistrationController := &controllers.TalkRegistrationController{}

	router.HandleFunc(talkRegistrationPath+"/{talkID:[0-9]+}/confirm", talkRegistrationController.ConfirmTalkRegistrationChange).Methods("GET")
	router.HandleFunc(talkRegistrationPath, talkRegistrationController.Create).Methods("POST")
	router.HandleFunc(talkRegistrationPath+"/{talkID:[0-9]+}", talkRegistrationController.GetTalkRegistrationsWithTalkID).Methods("GET")

	return router
}
