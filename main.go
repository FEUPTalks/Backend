package main

import (
	"log"
	"net/http"

	"github.com/RAyres23/back-end/routers"
)

const (
	listeningPort string = "8144"
)

func main() {
	router := routers.InitRoutes()
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":"+listeningPort, nil))
}
