package postgres

import (
	"database/sql"
	"fmt"
	"go-backed/app/configuration"
	"go-backed/app/types"
	"go-backed/app/types/errors"
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

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(config *configuration.Config) (*UserRepo, error) {
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

	return &UserRepo{
		DB: db,
	}, nil
}

func (ur *UserRepo) GetUserByID(id int) (types.User, error) {
	var user types.User
	query := `SELECT id, email, password FROM users WHERE id = $1`
	err := ur.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (ur *UserRepo) CreateUser(email, password string, isPremium bool) (*types.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := types.NewUser(email, string(hashedPassword), false)

	log.Println(user.Password)

	query := `INSERT INTO users(email, hashed_password, created_at, updated_at, is_premium)  VALUES($1, $2, $3, $4, $5) RETURNING id, email, created_at, updated_at, is_premium`
	err = ur.DB.QueryRow(query, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsPremium).Scan(
		&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsPremium)
	if err != nil {
		return nil, err
	}

	log.Println(user)
	return user, nil
}

func (ur *UserRepo) UpdateUser(user *types.User) error {
	query := `UPDATE users SET email = $1, updated_at = $3, is_premium = $4 WHERE email = $1`
	_, err := ur.DB.Exec(query, user.Email, user.UpdatedAt, user.IsPremium)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	query := `SELECT id, email, hashed_password FROM users WHERE email = $1`
	err := ur.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
