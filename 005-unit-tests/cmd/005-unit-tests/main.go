package main

import (
	"github.com/prokhorind/go_course/005-unit-tests/handler"
	"github.com/prokhorind/go_course/005-unit-tests/repo"
	"github.com/prokhorind/go_course/005-unit-tests/service"
)

func main() {
	r := repo.NewUserRepo()
	s := service.NewUserService(r)
	h := handler.NewUserHandler(s)

	h.HandleGetUser(1)
}
