package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

func WelcomeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome, " + c.GetString("username") + "!"})
	}
}
