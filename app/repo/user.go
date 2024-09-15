package repo

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	isPremium bool
}

func NewUser(id int, email, password string) *User {
	return &User{
		ID:        id,
		Email:     email,
		Password:  password,
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		isPremium: false,
	}
}

func NewAdminUser(id int, email string, password string) *User {
	return &User{
		ID:        id,
		Email:     email,
		Password:  password,
		IsAdmin:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		isPremium: true,
	}
}

type UserStore interface {
	GetUserByID(id int) (User, error)
}

type userStore struct {
	store Repository
}

func NewUserStore(store Repository) UserStore {
	return &userStore{
		store: store,
	}
}

func (us *userStore) GetUserByID(id int) (User, error) {
	return us.store.GetUserByID(id)
}
