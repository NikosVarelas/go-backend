package route

import (
	"go-backed/app/cache"
	"go-backed/app/configuration"
	"go-backed/app/handlers"
	"go-backed/app/repo"
	"go-backed/app/token"

	"github.com/gin-gonic/gin"
)

func NewHomeRouter(r *gin.Engine, userRepo repo.UserRepo, tokenMaker *token.JWTMaker, cache cache.Cache, config *configuration.Config) {
	// Public routes

	// Protected routes
	homeRoutes := r.Group("/")

	homeRoutes.GET("/", handlers.Home())

	// Webhook route
	r.POST("/webhook", handlers.Webhook(userRepo))
}
