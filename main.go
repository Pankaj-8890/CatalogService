package main

import (
	"catalogService/controller"
	"catalogService/middleware"
	"catalogService/model"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {

	dbClient := middleware.DatabaseConnection()
	serviceDb := model.NewServiceDb(dbClient)
	router := mux.NewRouter()
	controller.InitController(router,serviceDb)

	log.Printf("listening on port", http.ListenAndServe(":8082", router))
}
