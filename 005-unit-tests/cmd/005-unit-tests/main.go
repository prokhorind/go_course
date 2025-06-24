package main

import (
	"github.com/prokhorind/go_course/005-unit-tests/internal/handler"
	"github.com/prokhorind/go_course/005-unit-tests/internal/repo"
	"github.com/prokhorind/go_course/005-unit-tests/internal/service"
)

func main() {
	r := repo.NewUserRepo()
	s := service.NewUserService(r)
	h := handler.NewUserHandler(s)

	h.HandleGetUser(1)
}
