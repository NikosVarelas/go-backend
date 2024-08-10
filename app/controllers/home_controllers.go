package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func HandleHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Home Page",
			"Message": "Welcome to the home page",
		})
	}
}