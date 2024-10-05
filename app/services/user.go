package services

import (
	"go-backed/app/repo"
	"go-backed/app/types"

	bcrypt "golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db repo.UserRepo
}

func NewUserService(db repo.UserRepo) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(email, password string, isPremium bool) (*types.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := types.NewUser(email, string(hashedPassword), false)

	// Save user to database
	_, err = s.db.CreateUser(user)
	if err != nil {
		return nil, err
	}
	// ...

	return user, nil
}

func (s *UserService) GetUserByID(id int) (types.User, error) {
	return s.db.GetUserByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*types.User, error) {
	return s.db.GetUserByEmail(email)
}

func (s *UserService) UpdateUser(user *types.User) error {
	return s.db.UpdateUser(user)
}

func (s *UserService) Login(email, password string) (*types.User, error) {
	user, err := s.db.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
