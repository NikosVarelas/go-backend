package route

import (
	"go-backed/app/cache"
	"go-backed/app/configuration"
	"go-backed/app/controllers"
	"go-backed/app/middleware"
	"go-backed/app/store"
	"go-backed/app/token"
	"time"

	"github.com/gin-gonic/gin"
)

func NewHomeRouter(r *gin.Engine, repo store.Store, tokenMaker *token.JWTMaker, cache cache.Cache, config *configuration.Config) {
	// Public routes
	r.Use(middleware.RateLimitMiddleware(cache, config.RateLimit.MaxRequests, time.Duration(config.RateLimit.TimeInterval)*time.Minute))
	authRoutes := r.Group("/auth")
	authRoutes.GET("/login", controllers.LoginIndex())
	authRoutes.GET("/sign-up", controllers.SignUp())
	authRoutes.POST("/sign-up", controllers.SignUpSubmit(repo))
	authRoutes.POST("/login", controllers.LoginUser(repo, tokenMaker))
	authRoutes.GET("/logout", controllers.LogoutUser())

	// Protected routes
	protectedRoutes := r.Group("/")
	protectedRoutes.Use(middleware.AuthMiddleware(config))

	protectedRoutes.GET("/", controllers.Home())
}
