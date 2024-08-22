package middleware

import (
	"context"
	"fmt"
	"go-backed/app/token"

	"github.com/gin-gonic/gin"
)

type authKey struct{}

func AuthMiddleware(jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err:= verifyClaimsFromCookie(c, token.NewJWTMaker(jwtKey))
		if err != nil {
			c.Redirect(302, "/auth/login")
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), authKey{}, claims))
		c.Next()
	}
}

func verifyClaimsFromCookie(c *gin.Context, tokenMaker *token.JWTMaker) (*token.UserClaims, error) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		return nil, fmt.Errorf("missing access token")
	}
	claims, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token")
	}
	return claims, nil
}
