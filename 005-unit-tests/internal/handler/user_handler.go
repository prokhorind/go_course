package handler

import (
	"fmt"
	"github.com/prokhorind/go_course/005-unit-tests/internal/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
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
