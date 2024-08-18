package store

import "time"

type User struct {
	ID       int
	Email    string
	Password string
	IsAdmin  bool
	CreatedAt time.Time
	UpdateAt time.Time
}

func NewUser (id int, email, password string) *User {
	return &User{
		ID: id,
		Email: email,
		Password: password,
		IsAdmin: false,
		CreatedAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

func NewAdminUser(id int, email string, password string) *User {
	return &User{
		ID: id,
		Email: email,
		Password: password,
		IsAdmin: true,
		CreatedAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

type UserStore interface {
	GetUserByID(id int) (User, error)
}

type userStore struct {
	store Store
}

func NewUserStore(store Store) UserStore {
	return &userStore{
		store: store,

	}
}

func (us *userStore) GetUserByID(id int) (User, error) {
	return us.store.GetUserByID(id)}