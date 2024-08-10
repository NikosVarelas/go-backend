package models

import "database/sql"


var repo Database

type Models struct {
	User User
}

func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = NewPGDatabase(conn) // the function that actually hooks up mysql database to our models
	} else {
		repo = NewPGDatabase(nil)
	}

	return &Models{
		User: User{},
	}
}

func (m *Models) GetUserByID(id int) (User, error) {
	return repo.GetUserByID(id)
}