package routers

import (
	"github.com/FEUPTalks/Backend/controllers"
	"github.com/gorilla/mux"
)

const (
	talkRegistrationLogPath string = "/talkRegistrationLog"
)

//Sets the routes for /talkRegistration functionalities
func SetTalkRegistrationLogRoutes(router *mux.Router) *mux.Router {

	talkRegistrationLogController := &controllers.TalkRegistrationLogController{}

	router.HandleFunc(talkRegistrationLogPath+"/{talkID:[0-9]+}", talkRegistrationLogController.GetTalkRegistrationLogsWithTalkID).Methods("GET")

	return router
}
