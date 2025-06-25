package handler

import (
	"fmt"
	"github.com/prokhorind/go_course/005-unit-tests/internal/models"
)

type UserService interface {
	FetchUser(id int) (*models.User, error)
}

type UserHandler struct {
	svc UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{svc: s}
}

func (h *UserHandler) HandleGetUser(id int) {
	user, err := h.svc.FetchUser(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User: %+v\n", user)
}
