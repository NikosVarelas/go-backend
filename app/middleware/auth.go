package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAuthenticated := false

		if !isAuthenticated {
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
			return
		}

		c.Next()
	}
}