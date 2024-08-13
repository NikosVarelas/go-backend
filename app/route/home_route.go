package route

import (
	"go-backed/app/controllers"
	"go-backed/app/middleware"
	"go-backed/app/store"

	"github.com/gin-gonic/gin"
)

func NewHomeRouter(r *gin.Engine, repo store.Store) {
    // Public routes
	authRoutes := r.Group("/auth")
    authRoutes.GET("/login", controllers.LoginIndex())
    authRoutes.GET("/auth/sign-up", controllers.SignUp())
    authRoutes.POST("/sign-up", controllers.SignUpSubmit(repo))
    authRoutes.POST("/login", controllers.Login(repo))

    // Protected routes
    protectedRoutes := r.Group("/")
    protectedRoutes.Use(middleware.AuthMiddleware())

    protectedRoutes.GET("/", controllers.Home())
}

