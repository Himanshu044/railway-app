package routes

import (
	"backend/controllers"
	"backend/helpers"

	"github.com/gin-gonic/gin"
)

func SetupTrainRoutes(router *gin.Engine) {
	trainGroup := router.Group("/trains")
	{
		trainGroup.GET("/", helpers.RequireAuth, controllers.GetTrains)
	}
}
