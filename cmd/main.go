package main

import (
	"fmt"

	"github.com/hokkung/colonnello/entity"
	"github.com/hokkung/colonnello/repository"
	"github.com/hokkung/colonnello/service"
)

func main() {

	db, _, _ := entity.ProvideDB()

	userRepo := repository.ProvideUserRepository(db)

	userService := service.ProvideUserService(userRepo)

	fmt.Println(userService.GetAll())
	fmt.Println(userService.GetName("hok"))

}
