package types

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	IsPremium bool
}

func NewUser(email, password string, isPremium bool) *User {
	return &User{
		Email:     email,
		Password:  password,
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsPremium: isPremium,
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
		IsPremium: true,
	}
}
