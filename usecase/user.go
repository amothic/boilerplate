package usecase

import (
	"github.com/amothic/boilerplate/entity"
	"github.com/amothic/boilerplate/repository"
)

type UserInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(ur repository.UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: ur,
	}
}

func (interactor *UserInteractor) CreateUser(name string) (*entity.User, error) {
	user := entity.NewUser(name)
	err := interactor.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (interactor *UserInteractor) GetAll() ([]*entity.User, error) {
	return interactor.userRepository.GetAll()
}
