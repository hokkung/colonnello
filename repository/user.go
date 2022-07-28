package repository

import (
	"github.com/hokkung/colonnello/entity"
	"github.com/hokkung/colonnello/internal/repository"
	"gorm.io/gorm"
)

type UserRepository[T entity.Entity] interface {
	repository.Repository[T]
}

type userRepository[T entity.Entity] struct {
	repository.BaseRepository[T]
}

func ProvideUserRepository(db *gorm.DB) UserRepository[entity.User] {
	return &userRepository[entity.User]{
		BaseRepository: repository.BaseRepository[entity.User]{
			DB: db,
		},
	}
}
