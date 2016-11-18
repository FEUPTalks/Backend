package routers

import (
	"github.com/FEUPTalks/Backend/controllers"
	"github.com/gorilla/mux"
)

const (
	picturePath string = "/picture"
)

//SetPictureRoutes sets the routes for /file functionalities
func SetPictureRoutes(router *mux.Router) *mux.Router {

	pictureController := &controllers.PictureController{}

	router.HandleFunc(picturePath+"/{pictureID:[0-9]+}", pictureController.Download).Methods("GET")
	router.HandleFunc(picturePath, pictureController.Upload).Methods("POST")

	return router
}
