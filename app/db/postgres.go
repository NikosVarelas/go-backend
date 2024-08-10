package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	maxOpenDbConn = 25
	maxIdleDBConn = 25
	maxDBLifetime = 5 * time.Minute
)

func InitPostgres() (*sql.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// test our database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetConnMaxIdleTime(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)

	fmt.Println("Connected to postgres db")
	return db, nil
}
