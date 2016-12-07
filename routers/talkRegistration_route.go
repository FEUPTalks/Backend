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

	router.HandleFunc(talkRegistrationPath, talkRegistrationController.Create).Methods("POST")

	return router
}
