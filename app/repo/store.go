package repo

import (
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Store interface {
	GetUserByID(id int) (User, error)
	CreateNewUser(email, password string) (User, error)
	GetUserByEmail(email string) (User, error)
}
