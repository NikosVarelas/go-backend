package repo

import (
	"go-backed/app/types"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type UserRepo interface {
	GetUserByID(id int) (types.User, error)
	CreateNewUser(email, password string, isPremium bool) (*types.User, error)
	GetUserByEmail(email string) (*types.User, error)
}
