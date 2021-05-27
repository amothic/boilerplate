package repository

import "github.com/amothic/boilerplate/entity"

type UserRepository interface {
	GetAll() ([]*entity.User, error)
	Create(user *entity.User) error
}
