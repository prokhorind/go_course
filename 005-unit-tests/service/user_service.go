package service

import (
	"errors"
	"github.com/prokhorind/go_course/005-unit-tests/models"
	"strings"
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
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Business rule: restrict access to admin users with ID < 10
	if strings.HasPrefix(user.Name, "Admin") && user.ID < 10 {
		return nil, errors.New("access to admin users is restricted")
	}

	return user, nil
}
