package main

import (
	"log"
	"net/http"

	"github.com/RAyres23/LESTeamB-backend/database"
	"github.com/RAyres23/LESTeamB-backend/routers"
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
