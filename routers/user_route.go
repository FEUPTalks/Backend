package routers

import (
	"github.com/FEUPTalks/Backend/controllers"
	"github.com/gorilla/mux"
)

const (
	userPath string = "/user"
)

//SetUserRoutes sets the routes for /file functionalities
func SetUserRoutes(router *mux.Router) *mux.Router {

	userController := &controllers.UserController{}

	router.HandleFunc(userPath+"/{userID:[0-9a-z]+}", userController.GetUser).Methods("GET")
	router.HandleFunc(userPath+"/{userID:[0-9a-z]+}", userController.SetUser).Methods("POST")
	router.HandleFunc(userPath, userController.Create).Methods("POST")

	return router
}
