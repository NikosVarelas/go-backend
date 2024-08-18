package route

import (
	"go-backed/app/controllers"
	"go-backed/app/store"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuthRouter(timeout time.Duration, db *store.Store, gin *gin.Engine) {
	AuthRouter := gin.Group("/auth")
	// All Public APIs
	AuthRouter.POST("/", controllers.LoginIndex())
}