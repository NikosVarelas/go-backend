package route

import (
	"go-backed/app/store"
	"go-backed/app/token"

	"github.com/gin-gonic/gin"
)

func NewRouter(repo store.Store, tokenMaker *token.JWTMaker) *gin.Engine {
	router := gin.Default()
	NewHomeRouter(router, repo, tokenMaker)
	return router
}
