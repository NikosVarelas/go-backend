package handlers

import (
	"go-backed/app/services"
	"go-backed/app/token"
	"go-backed/templates"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginUser(us *services.UserService, tokenMaker *token.JWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u LoginUserReq
		if err := c.ShouldBind(&u); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		email := u.Email
		password := u.Password

		user, err := us.Login(email, password)
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

func SignUpSubmit(us *services.UserService) gin.HandlerFunc {
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

		_, err := us.CreateUser(email, password, false)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to create user")
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Set("Content-Type", "text/html")

		c.Redirect(http.StatusFound, "/")
	}
}
