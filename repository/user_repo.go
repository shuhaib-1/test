package repository

import (
	"fmt"
	"sync"
	"test/domain"
)

type UserRepository interface {
	Save(user domain.User) error
	FindByID(id int) (domain.User, error)
}

type UserRepositoryImpl struct {
	data map[string]interface{}
	mu   sync.Mutex // Add a mutex to protect the map
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		data: make(map[string]interface{}),
	}
}

func (r *UserRepositoryImpl) Save(user domain.User) error {
    r.mu.Lock()         // Lock the mutex before modifying the map
    defer r.mu.Unlock() // Unlock the mutex after modification

    // Assuming the user has an ID field of type int
    key := fmt.Sprintf("%d", user.ID)
    r.data[key] = user
    return nil
}

func (r *UserRepositoryImpl) Get(key string) (interface{}, bool) {
	r.mu.Lock()         // Lock the mutex before reading the map
	defer r.mu.Unlock() // Unlock the mutex after reading
	val, ok := r.data[key]
	return val, ok
}

func (r *UserRepositoryImpl) FindByID(id int) (domain.User, error) {
	r.mu.Lock()         // Lock the mutex before accessing the map
	defer r.mu.Unlock() // Unlock the mutex after accessing

	// Assuming the key is stored as a string representation of the ID
	key := fmt.Sprintf("%d", id)
	if val, ok := r.data[key]; ok {
		// Type assert the value to domain.User
		if user, ok := val.(domain.User); ok {
			return user, nil
		}
		return domain.User{}, fmt.Errorf("data type mismatch for ID %d", id)
	}
	return domain.User{}, fmt.Errorf("user with ID %d not found", id)
}
