package main

import (
	"go-backed/app/route"
	"go-backed/app/store"
	"go-backed/app/token"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main () {
	dbConfig := &store.PGConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB_NAME"),
	}
	db, err:= store.NewPGStore(dbConfig)

	if err != nil {
		log.Fatal(err)
	}
	tokenMaker := token.NewJWTMaker(os.Getenv("JWT_SECRET_KEY"))
	router := route.NewRouter(db, tokenMaker)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDR")

	log.Println("Server started on", listenAddr)

	server := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}

	userStore := store.NewUserStore(db)
	user, _ := userStore.GetUserByID(1)

	log.Println("getting user by id")
	log.Println(user.Email)

	server.ListenAndServe()
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}