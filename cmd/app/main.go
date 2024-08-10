package main

import (
	"go-backed/app/configuration"
	"go-backed/app/db"
	"go-backed/app/route"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type app struct {
	Application *configuration.Application
}

func main () {
	app := &app{}
	router := route.NewRouter()
	router.LoadHTMLGlob("app/views/*")

	listenAddr := os.Getenv("HTTP_LISTEN_ADDR")

	log.Println("Server started on", listenAddr)

	server := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}
	db, err:= db.InitPostgres()

	if err != nil {
		log.Fatal(err)
	}

	app.Application = configuration.New(db)

	server.ListenAndServe()
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}