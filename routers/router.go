package routers

import "github.com/gorilla/mux"

// InitRoutes Initializes the System routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetTalksRoutes(router)
	router = SetPictureRoutes(router)
	return router
}
