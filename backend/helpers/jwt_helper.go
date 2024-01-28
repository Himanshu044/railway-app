package helpers

import (
	"backend/models"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(user_id uint) (map[string]string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("Jwt_secret")))

	if err != nil {
		return nil, errors.New("Failed to create token")
	}

	return map[string]string{
		"token": tokenString,
	}, nil
}

func RequireAuth(context *gin.Context) {
	tokenString, err := context.Cookie("Authorization")

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Authenicate user.",
		})
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("Jwt_secret")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		DB := GetDbInstance()
		var user models.User
		DB.First(&user, claims["sub"])

		if user.ID == 0 {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		context.Set("user", user)
		context.Next()
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}
