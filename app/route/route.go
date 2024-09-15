package route

import (
	"go-backed/app/cache"
	"go-backed/app/configuration"
	"go-backed/app/repo/postgres"
	"go-backed/app/token"
	"log"

	"github.com/gin-gonic/gin"
)

func NewRouter(config *configuration.Config) *gin.Engine {
	db, err := postgres.NewUserRepo(config)
	if err != nil {
		log.Fatal(err)
	}

	redis := cache.NewRedisCache()

	ping, err := redis.Ping()

	log.Println("Conneting to redis", ping)

	if err != nil {
		log.Println(err)
	}

	tokenMaker := token.NewJWTMaker(config)

	router := gin.Default()
	NewHomeRouter(router, db, tokenMaker, redis, config)
	NewAuthRoute(router, db, tokenMaker, redis, config)
	return router
}
