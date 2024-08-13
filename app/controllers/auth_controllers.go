package controllers

import (
	"context"
	"go-backed/app/store"
	"go-backed/templates"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(repo store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")

		user, err := repo.GetUserByEmail(email)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get user")
			return
		}
		log.Println(password)
		log.Println(user)
		pwdMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if pwdMatch != nil {
			c.String(http.StatusUnauthorized, "Invalid credentials")
			return
		}
		c.Redirect(http.StatusFound, "/")
	}
}

func LoginIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the status and content type
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Set("Content-Type", "text/html")

		// Render the templ component and write it to the response
		err := templates.Index().Render(context.Background(),c.Writer)
		if err != nil {
			// Handle the error
			c.String(http.StatusInternalServerError, "Failed to render template")
			return
		}
	}
}

func SignUpSubmit (repo store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")
		confirmPassowrd := c.Request.FormValue("confirm-password")


		if email == "" || password == "" || confirmPassowrd == "" {
			c.String(http.StatusBadRequest, "email, password and confirm password are required")
			return
		}

		if password != confirmPassowrd {
			c.String(http.StatusBadRequest, "passwords do not match")
		}

		_, err := repo.CreateNewUser(email, password)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to create user")
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Set("Content-Type", "text/html")
		
		c.Redirect(http.StatusFound, "/protected/home")
	}
}