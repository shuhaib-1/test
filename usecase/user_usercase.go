package usecase

import (
	"test/domain"
	"test/repository"
)

type UserUseCase interface {
	CreateUser(user domain.User) error
	GetUserById(id int) (domain.User, error)
}

type UserUseCaseImpt struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &UserUseCaseImpt{repo: repo}
}

func (uc *UserUseCaseImpt) CreateUser(user domain.User) error {
	return uc.repo.Save(user)
}

func (uc *UserUseCaseImpt) GetUserById(id int) (domain.User, error) {
	return uc.repo.FindByID(id)
}
