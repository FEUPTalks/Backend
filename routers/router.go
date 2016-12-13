package routers

import "github.com/gorilla/mux"

// InitRoutes Initializes the System routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetAuthenticationRoutes(router)
	router = SetTalksRoutes(router)
	router = SetPictureRoutes(router)
	router = SetTalkRegistrationRoutes(router)
	router = SetTalkRegistrationLogRoutes(router)

	return router
}
