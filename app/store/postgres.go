package store

import (
	"database/sql"
	"fmt"
	"go-backed/app/configuration"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
	"golang.org/x/crypto/bcrypt"
)

const (
	maxOpenDbConn = 25
	maxIdleDBConn = 25
	maxDBLifetime = 5 * time.Minute
)

type PGStore struct {
	DB *sql.DB
}

func NewPGStore(config *configuration.Config) (*PGStore, error) {
	dbConfig := config.Database
	dsn := fmt.Sprintf("postgres://%s:%s@%s%s/%s?sslmode=disable", dbConfig.PostgresUser, dbConfig.PostgresPassword, dbConfig.PostgresHost, dbConfig.PostgresPort, dbConfig.PostgresName)
	db, err := sql.Open("postgres", dsn)
	log.Println(dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetConnMaxIdleTime(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)
	fmt.Println("Connected to postgres db")

	return &PGStore{
		DB: db,
	}, nil
}

func (pg *PGStore) GetUserByID(id int) (User, error) {
	var user User
	query := `SELECT id, email, password FROM users WHERE id = $1`
	err := pg.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (pg *PGStore) CreateNewUser(email, password string) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user := NewUser(0, email, string(hashedPassword))

	log.Println(user.Password)

	query := `INSERT INTO users(email, password) VALUES($1, $2) RETURNING id, email, password, created_at, updated_at`
	err = pg.DB.QueryRow(query, email, hashedPassword).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdateAt)
	if err != nil {
		return User{}, err
	}

	log.Println(user)
	return *user, nil
}

func (pg *PGStore) GetUserByEmail(email string) (User, error) {
	var user User
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := pg.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
