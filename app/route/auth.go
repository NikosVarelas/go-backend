package route

import (
	"go-backed/app/cache"
	"go-backed/app/configuration"
	"go-backed/app/handlers"
	"go-backed/app/middleware"
	"go-backed/app/repo"
	"go-backed/app/token"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuthRoute(r *gin.Engine, repo repo.UserRepo, tokenMaker *token.JWTMaker, cache cache.Cache, config *configuration.Config) {
	// Public routes
	r.Use(middleware.RateLimitMiddleware(cache, config.RateLimit.MaxRequests, time.Duration(config.RateLimit.TimeInterval)*time.Minute))
	authRoutes := r.Group("/auth")
	authRoutes.GET("/login", handlers.LoginIndex())
	authRoutes.GET("/sign-up", handlers.SignUp())
	authRoutes.POST("/sign-up", handlers.SignUpSubmit(repo))
	authRoutes.POST("/login", handlers.LoginUser(repo, tokenMaker))
	authRoutes.GET("/logout", handlers.LogoutUser())

}
