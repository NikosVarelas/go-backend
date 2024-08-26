package route

import (
	"go-backed/app/cache"
	"go-backed/app/configuration"
	"go-backed/app/store"
	"go-backed/app/token"

	"github.com/gin-gonic/gin"
)

func NewRouter(repo store.Store, tokenMaker *token.JWTMaker, cache cache.Cache, config *configuration.Config) *gin.Engine {
	router := gin.Default()
	NewHomeRouter(router, repo, tokenMaker, cache, config)
	return router
}
