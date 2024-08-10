package route

import (
	"github.com/gin-gonic/gin"
)

func NewRouter () * gin.Engine{
	router := gin.Default()
	NewHomeRouter(router)
	return router
}