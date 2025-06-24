package repo

import (
	"github.com/prokhorind/go_course/005-unit-tests/internal/models"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetUserByID(id int) (*models.User, error) {
	// Simulate DB call
	return &models.User{ID: id, Name: "Alice"}, nil
}
