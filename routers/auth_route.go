package routers

import (
	"github.com/FEUPTalks/Backend/controllers"
	"github.com/FEUPTalks/Backend/core/authentication"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetAuthenticationRoutes sets authentication routes
func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")
	return router
}
