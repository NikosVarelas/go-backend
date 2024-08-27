package main

import (
	"go-backed/app/configuration"
	"go-backed/app/route"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	config, err := configuration.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	router := route.NewRouter(config)
	router.Static("/static/", "./static")

	log.Println("Server started on", config.Server.HTTPListenAddr)

	server := &http.Server{
		Addr:    config.Server.HTTPListenAddr,
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
