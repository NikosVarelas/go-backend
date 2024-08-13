package route

import (
	"go-backed/app/controllers"
	"go-backed/app/models"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuthRouter(timeout time.Duration, db *models.Database, gin *gin.Engine) {
	AuthRouter := gin.Group("/auth")
	// All Public APIs
	AuthRouter.POST("/", controllers.LoginIndex())
}