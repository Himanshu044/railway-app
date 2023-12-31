package main

import (
	"backend/helpers"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	helpers.LoadEnvVariables()
	helpers.SyncDatabase()
	router.Run()
}
