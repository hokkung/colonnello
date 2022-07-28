package service

import (
	"context"

	"github.com/hokkung/colonnello/entity"
	"github.com/hokkung/colonnello/repository"
)

type UserService interface {
	GetAll() []entity.User
	GetName(name string) entity.User
}

type userService struct {
	userRepository repository.UserRepository[entity.User]
}

func ProvideUserService(userRepository repository.UserRepository[entity.User]) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) GetAll() []entity.User {
	users, err := u.userRepository.FindAll()
	if err != nil {
		panic(err)
	}

	return users
}

func (u userService) GetName(name string) entity.User {
	user, _, err := u.userRepository.FindByName(context.Background(), name)
	if err != nil {
		panic(err)
	}

	return user
}
