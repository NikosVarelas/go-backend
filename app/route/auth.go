package route

import (
	"go-backed/app/cache"
	"go-backed/app/configuration"
	"go-backed/app/handlers"
	"go-backed/app/middleware"
	"go-backed/app/services"
	"go-backed/app/token"

	"time"

	"github.com/gin-gonic/gin"
)

func NewAuthRoute(r *gin.Engine, tokenMaker *token.JWTMaker, cache cache.Cache, config *configuration.Config, userService *services.UserService) {
	// Public routes
	r.Use(middleware.RateLimitMiddleware(cache, config.RateLimit.MaxRequests, time.Duration(config.RateLimit.TimeInterval)*time.Minute))
	authRoutes := r.Group("/auth")
	authRoutes.GET("/login", handlers.LoginIndex())
	authRoutes.GET("/sign-up", handlers.SignUp())
	authRoutes.POST("/sign-up", handlers.SignUpSubmit(userService))
	authRoutes.POST("/login", handlers.LoginUser(userService, tokenMaker))
	authRoutes.GET("/logout", handlers.LogoutUser())

}
