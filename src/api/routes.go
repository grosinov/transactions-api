package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grosinov/transactions-api/src/api/controllers"
)

func SetupRouter(controller *controllers.TransactionsController) *gin.Engine {
	router := gin.Default()

	router.POST("/migrate", controller.MigrateTransactions)
	router.GET("/users/:user_id/balance", controller.GetBalance)

	return router
}
