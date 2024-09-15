package controllers

import (
	"go-backed/app/repo"
	"go-backed/app/token"
	"go-backed/templates"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(repo repo.Store, tokenMaker *token.JWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u LoginUserReq
		if err := c.ShouldBind(&u); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		email := u.Email
		password := u.Password

		user, err := repo.GetUserByEmail(email)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get user")
			return
		}

		pwdMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if pwdMatch != nil {
			c.String(http.StatusUnauthorized, "Invalid credentials")
			return
		}

		accessToken, _, err := tokenMaker.CreateToken(user.ID, user.Email, user.IsAdmin)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to create access token")
			return
		}
		cookie := http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Expires:  time.Now().Add(tokenMaker.GetExpiration()),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Secure:   true,
			Path:     "/",
		}
		http.SetCookie(c.Writer, &cookie)

		c.Redirect(http.StatusFound, "/")
	}
}

func LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("access_token", "", -1, "/", "localhost:3000", false, true)
		c.SetCookie("refresh_token", "", -1, "/", "localhost:3000", false, true)
		c.Redirect(http.StatusFound, "/auth/login")
	}
}

func LoginIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the status and content type
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Set("Content-Type", "text/html")
		// Render the templ component and write it to the response
		t := templates.Login()
		err := templates.Layout(t, "login", false).Render(c.Request.Context(), c.Writer)
		if err != nil {
			// Handle the error
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func SignUpSubmit(repo repo.Store) gin.HandlerFunc {
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

		c.Redirect(http.StatusFound, "/")
	}
}
