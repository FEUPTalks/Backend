package routers

import (
	"github.com/FEUPTalks/Backend/controllers"
	"github.com/FEUPTalks/Backend/core/authentication"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const (
	talkPath string = "/talks"
)

//SetTalksRoutes sets the routes for /talks functionalities
func SetTalksRoutes(router *mux.Router) *mux.Router {

	talkController := &controllers.TalkController{}

	router.HandleFunc(talkPath, talkController.Index).Methods("GET")
	router.Handle(talkPath+"/all",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(talkController.All),
		)).Methods("GET")
	router.HandleFunc(talkPath+"/{talkID:[0-9]+}", talkController.GetTalk).Methods("GET")
	router.Handle(talkPath+"/{talkID:[0-9]+}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(talkController.SetTalk),
		)).Methods("PUT")
	router.Handle(talkPath+"/{talkID:[0-9]+}/SetState",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(talkController.SetTalkState),
		)).Methods("PUT")
	router.HandleFunc(talkPath, talkController.Create).Methods("POST")

	return router
}
