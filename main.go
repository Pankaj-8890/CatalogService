package main

import (
	"catalogService/controller"
	"catalogService/middleware"
	"catalogService/model"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {


	dbClient := middleware.DatabaseConnection()
	serviceDb := model.NewServiceDb(dbClient)

	router := mux.NewRouter()
	controller.InitController(router,serviceDb)

	if err := http.ListenAndServe(":8082", router); err != nil {
		fmt.Println("Error while creating the server listen: ", err)
	}
}
