package main

import (
	"DI/controllers"
	"DI/ds"
	"DI/services"
	"DI/utils"
	"net/http"
)

func main() {
	log := utils.LoggerAdapter(utils.LogOutput)
	dataStore := ds.NewSimpleDataStore()
	logic := services.NewSimpleLogic(log, dataStore)
	controller := controllers.NewController(log, logic)
	http.HandleFunc("/hello", controller.SayHello)
	http.ListenAndServe(":8080", nil)

}
