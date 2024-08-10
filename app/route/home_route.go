package route

import (
	"go-backed/app/controllers"

	"github.com/gin-gonic/gin"
)

func NewHomeRouter(r *gin.Engine) {
	r.GET("/", controllers.HandleHome())
}
