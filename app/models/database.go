package models

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Database interface {
	GetUserByID(id int) (User, error)
}

type PGDatabase struct {
	DB *sql.DB
}

func NewPGDatabase(conn *sql.DB) *PGDatabase {
	return &PGDatabase{
		DB: conn,
	}
}