package main

import (
	"log"
	"net/http"

	"github.com/FEUPTalks/Backend/core/authentication"
	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/routers"
	"github.com/FEUPTalks/Backend/settings"
	"github.com/urfave/negroni"
	"github.com/rs/cors"
)

const (
	listeningPort string = "8144"
)

func main() {

	//Init Settings
	settings.Init()

	//Init Keys
	_, err := authentication.GetJWTAuthenticationBackend()
	if err != nil {
		log.Fatal(err)
		return
	}

	//Init talk database access
	talkInstance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer talkInstance.CloseConnection()

	err = talkInstance.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	//Init user database access
	userInstance, err := database.GetUserDatabaseManagerInstance()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer userInstance.CloseConnection()

	err = userInstance.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":"+listeningPort, n))
}
