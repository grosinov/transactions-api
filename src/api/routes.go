package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grosinov/transactions-api/docs"
	"github.com/grosinov/transactions-api/src/api/controllers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(controller *controllers.TransactionsController) *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	v1 := router.Group("/api/v1")

	v1.POST("/migrate", controller.MigrateTransactions)
	v1.GET("/users/:user_id/balance", controller.GetBalance)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
