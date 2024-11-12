package main

import (
	"github.com/grosinov/transactions-api/src/api/config"
	"github.com/grosinov/transactions-api/src/api/controllers"
	"github.com/grosinov/transactions-api/src/api/repositories"
	"github.com/grosinov/transactions-api/src/api/services"
	"log"
)

// @title Transaction API
// @version 1.0
// @description API to handle transactions
// @host localhost:8080
// @BasePath /api/v1
func main() {
	repository := repositories.NewTransactionsRepository(config.ConnectDatabase())
	service := services.NewTransactionsServiceImpl(repository)
	controller := controllers.NewTransactionsController(service)
	router := SetupRouter(controller)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
