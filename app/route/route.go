package route

import (
	"go-backed/app/cache"
	"go-backed/app/configuration"
	"go-backed/app/repo/postgres"
	"go-backed/app/services"
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

	userService := services.NewUserService(db)

	router := gin.Default()
	NewHomeRouter(router, userService)
	NewAuthRoute(router, tokenMaker, redis, config, userService)
	return router
}
