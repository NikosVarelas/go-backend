package user

import (
	"go-backed/app/models"
)

type Service interface {
	GetUserByID(id int) (models.User, error)
}

type service struct {
	Models *models.Models
}

func (s* service) GetUserByID(id int) (models.User, error) {
	return s.Models.GetUserByID(id)
}	

func NewService(models *models.Models) Service {
	return &service{
		Models: models,
	}
}