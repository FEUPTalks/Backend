package main

import (
	"log"
	"net/http"

	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/routers"
)

const (
	listeningPort string = "8144"
)

func main() {

	//Init database access
	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer instance.CloseConnection()

	err = instance.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	router := routers.InitRoutes()
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":"+listeningPort, nil))
}
