package route

import (
	"go-backed/app/cache"
	"go-backed/app/configuration"
	"go-backed/app/controllers"
	"go-backed/app/store"
	"go-backed/app/token"

	"github.com/gin-gonic/gin"
)

func NewHomeRouter(r *gin.Engine, repo store.Store, tokenMaker *token.JWTMaker, cache cache.Cache, config *configuration.Config) {
	// Public routes

	// Protected routes
	homeRoutes := r.Group("/")

	homeRoutes.GET("/", controllers.Home())

	// Webhook route
	r.POST("/webhook", controllers.Webhook())
}
