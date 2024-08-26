package main

import (
	"go-backed/app/cache"
	"go-backed/app/configuration"
	"go-backed/app/route"
	"go-backed/app/store"
	"go-backed/app/token"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	config := configuration.NewConfig()
	db, err := store.NewPGStore(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	redis := cache.NewRedisCache()

	ping, err := redis.Ping()

	log.Println("Conneting to redis", ping)

	if err != nil {
		log.Println(err)
	}
	tokenMaker := token.NewJWTMaker(os.Getenv("JWT_SECRET_KEY"))
	router := route.NewRouter(db, tokenMaker, redis, config)
	router.Static("/static/", "./static")

	listenAddr := os.Getenv("HTTP_LISTEN_ADDR")

	log.Println("Server started on", listenAddr)

	server := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}

	server.ListenAndServe()
}

func init() {
	log.Println(os.Getenv("IS_DOCKER"))
	log.Println(os.Getenv("POSTGRES_HOST"))
	if os.Getenv("IS_DOCKER") == "true" {
		return
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
