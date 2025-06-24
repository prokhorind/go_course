package service

import (
	"github.com/prokhorind/go_course/005-unit-tests/models"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
}

type UserService interface {
	FetchUser(id int) (*models.User, error)
}

type userServiceImpl struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userServiceImpl{repo: r}
}

func (s *userServiceImpl) FetchUser(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
