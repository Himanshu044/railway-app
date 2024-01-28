package controllers

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(context *gin.Context) {
	var user models.User
	if context.Bind(&user) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}
	user.Password = string(hash)
	DB := helpers.GetDbInstance()
	result := DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create user.",
			"message": result.Error,
		})
		return
	}
	context.JSON(http.StatusOK, user)
}

func Login(context *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	var user models.User
	DB := helpers.GetDbInstance()
	DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	tokenMap, err := helpers.GenerateToken(user.ID)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", tokenMap["token"], 3600*24*30, "", "", false, true)
	context.Header("AuthToken", tokenMap["token"])
	context.JSON(http.StatusOK, gin.H{"status": "Successfully logedin.."})
}

func Logout(context *gin.Context) {
	context.SetCookie("Authorization", "", 0, "/", "", false, true)
	context.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
