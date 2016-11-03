package routers

import (
	"github.com/RAyres23/back-end/controllers"
	"github.com/gorilla/mux"
)

const (
	talkPath string = "/talks"
)

// SetTalksRoutes sets the routes for /talks functionalities
func SetTalksRoutes(router *mux.Router) *mux.Router {

	talkController := &controllers.TalkController{}

	router.HandleFunc(talkPath, talkController.Index).Methods("GET")
	router.HandleFunc(talkPath, talkController.Create).Methods("POST")

	return router
}
