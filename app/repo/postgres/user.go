package postgres

import (
	"database/sql"
	"fmt"
	"go-backed/app/configuration"
	"go-backed/app/types"
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

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(config *configuration.Config) (*UserStore, error) {
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

	return &UserStore{
		DB: db,
	}, nil
}

func (us *UserStore) GetUserByID(id int) (types.User, error) {
	var user types.User
	query := `SELECT id, email, password FROM users WHERE id = $1`
	err := us.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (us *UserStore) CreateNewUser(email, password string, isPremium bool) (*types.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := types.NewUser(email, string(hashedPassword), isPremium)

	log.Println(user.Password)

	query := `INSERT INTO users(email, hashed_password, created_at, updated_at, is_premium)  VALUES($1, $2, $3, $4, $5) RETURNING id, email, created_at, updated_at, is_premium`
	err = us.DB.QueryRow(query, &user.Email, &user.Password, &user.IsPremium, &user.CreatedAt, &user.UpdatedAt).Scan(
		&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsPremium)
	if err != nil {
		return nil, err
	}

	log.Println(user)
	return user, nil
}

func (us *UserStore) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	query := `SELECT id, email, hashed_password FROM users WHERE email = $1`
	err := us.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
