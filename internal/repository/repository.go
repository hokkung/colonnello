package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository[T Entity] interface {
	FindAll() ([]T, error)
	FindByKey(ctx context.Context, id interface{}) (T, bool, error)
	FindByName(ctx context.Context, name string) (T, bool, error)
}

type BaseRepository[T Entity] struct {
	DB    *gorm.DB
	Model string
}

func (r BaseRepository[T]) FindAll() ([]T, error) {
	var ents []T
	res := r.DB.Table(r.Model).Find(&ents)
	if res.Error != nil {
		return ents, res.Error
	}

	return ents, nil
}

func (r BaseRepository[T]) FindByKey(ctx context.Context, id interface{}) (T, bool, error) {
	var ent T
	res := r.DB.Table(r.Model).Where("id = ?", id).First(&ent)
	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ent, false, res.Error
		}

		return ent, false, nil
	}

	return ent, true, nil
}

func (r BaseRepository[T]) FindByName(ctx context.Context, name string) (T, bool, error) {
	var ent T
	res := r.DB.Table(r.Model).Where("name = ?", name).First(&ent)
	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ent, false, res.Error
		}

		return ent, false, nil
	}

	return ent, true, nil
}
