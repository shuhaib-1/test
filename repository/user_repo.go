package repository

import (
	"errors"
	"test/domain"
)

type UserRepository interface {
	Save(user domain.User) error
	FindByID(id int) (domain.User, error)
}

type UserRepositoryImpl struct {
	users map[int]domain.User
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: make(map[int]domain.User),
	}
}

func (repo *UserRepositoryImpl) Save(user domain.User) error {

	if _, exists := repo.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	repo.users[user.ID] = user
	return nil
}

func (repo *UserRepositoryImpl) FindByID(id int) (domain.User, error) {

	user, exists := repo.users[id]

	if !exists {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}
