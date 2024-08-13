package controllers

import (
	"context"
	"go-backed/templates"
	"net/http"

	"github.com/gin-gonic/gin"
)



func SignUp () gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Set("Content-Type", "text/html")
		err := templates.SignUp().Render(context.Background(), c.Writer)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to render template")
			return
		}
	}
}

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Set("Content-Type", "text/html")
		err := templates.Home().Render(context.Background(), c.Writer)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to render template")
			return
		}
	}
}
