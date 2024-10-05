package route

import (
	"go-backed/app/handlers"
	"go-backed/app/services"

	"github.com/gin-gonic/gin"
)

func NewHomeRouter(r *gin.Engine, userService *services.UserService) {
	// Public routes

	// Protected routes
	homeRoutes := r.Group("/")

	homeRoutes.GET("/", handlers.Home())

	// Webhook route
	r.POST("/webhook", handlers.Webhook(userService))
}
