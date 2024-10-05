package repo

import (
	"go-backed/app/types"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type UserRepo interface {
	GetUserByID(id int) (types.User, error)
	CreateUser(user *types.User) (*types.User, error)
	UpdateUser(user *types.User) error
	GetUserByEmail(email string) (*types.User, error)
}
