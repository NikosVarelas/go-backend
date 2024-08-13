package route

import (
	"go-backed/app/store"

	"github.com/gin-gonic/gin"
)

func NewRouter (repo store.Store) * gin.Engine{
	router := gin.Default()
	NewHomeRouter(router, repo)
	return router
}