package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stepann0/balance-microservice/controllers"
	"github.com/stepann0/balance-microservice/models"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	models.ConnectDatabase()

	// Routes
	r.PUT("/", controllers.CreateAccount)
	r.GET("/:id", controllers.GetAccount)
	r.POST("/increase", controllers.IncreaseBalance)
	r.POST("/reserve", controllers.ReserveBalance)
	r.PUT("/accept", controllers.AcceptPayment)
	r.PUT("/decline", controllers.DeclinePayment)
	r.GET("/report/:month/:year", controllers.Report)
	r.DELETE("/:id", controllers.DeleteAccount)

	r.Run()
}
