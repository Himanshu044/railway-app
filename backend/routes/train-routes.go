package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTrainRoutes(router *gin.Engine) {
	trainGroup := router.Group("/trains")
	{
		trainGroup.GET("/", controllers.GetTrains)
	}
}
