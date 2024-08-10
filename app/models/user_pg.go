package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (pg *PGDatabase) GetUserByID(id int) (User, error) {
	var user User
	err := pg.DB.QueryRow("SELECT id, first_name, last_name, email, password FROM users WHERE id = $1", id).
	Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

