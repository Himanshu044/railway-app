package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupSessionRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
}
