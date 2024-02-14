package handlers

import (
	"userdata-redis-example-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", LoginHandler)
	}

	user := router.Group("/user")
	{
		user.POST("/register", RegisterHandler)
	}
	admin := router.Group("/admin")
	{
		admin.GET("/welcome", middleware.AuthMiddleware(), WelcomeHandler())
	}

}
