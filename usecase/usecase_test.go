package usecase_test

import (
	"fmt"
	"test/domain"
	"test/usecase"
	"testing"
)

type MockRepository struct {
	SaveFunc     func(user domain.User) error
	FindByIDFunc func(id int) (domain.User, error)
}

func (m *MockRepository) Save(user domain.User) error {
	return m.SaveFunc(user)
}

func (m *MockRepository) FindByID(id int) (domain.User, error) {
	return m.FindByIDFunc(id)
}

func TestUserUseCase(t *testing.T) {
	t.Run("CreateUser - valid input", func(t *testing.T) {
		mockRepo := &MockRepository{
			SaveFunc: func(user domain.User) error {
				if user.Name == "" || user.Email == "" {
					t.Errorf("save called with invalid user: %+v", user)
				}
				return nil
			},
		}

		uc := usecase.NewUserUseCase(mockRepo)
		user := domain.User{Name: "shuhaib", Email: "shuhaibpa@gmail.com"}
		err := uc.CreateUser(user)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("GetUserById - valid ID", func(t *testing.T) {
		mockRepo := &MockRepository{
			FindByIDFunc: func(id int) (domain.User, error) {
				if id == 1 {
					return domain.User{ID: 1, Name: "shuhaib", Email: "shuhaibpa85@gmail.com"}, nil
				}
				return domain.User{}, fmt.Errorf("user not found")
			},
		}

		uc := usecase.NewUserUseCase(mockRepo)

		user, err := uc.GetUserById(1)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if user.ID != 1 || user.Name != "shuhaib" || user.Email != "shuhaibpa85@gmail.com" {
			t.Errorf("unexpected user: %+v", user)
		}
	})

	t.Run("GetUserById - user not found", func(t *testing.T) {
		mockRepo := &MockRepository{
			FindByIDFunc: func(id int) (domain.User, error) {
				return domain.User{}, fmt.Errorf("user not found")
			},
		}

		uc := usecase.NewUserUseCase(mockRepo)
		_, err := uc.GetUserById(99)

		if err == nil {
			t.Errorf("expected error for non-existing user, got nil")
		}
	})

	t.Run("GetUserById - repository error", func(t *testing.T) {
		mockRepo := &MockRepository{
			FindByIDFunc: func(id int) (domain.User, error) {
				return domain.User{}, fmt.Errorf("database error")
			},
		}

		uc := usecase.NewUserUseCase(mockRepo)
		_, err := uc.GetUserById(1)

		if err == nil {
			t.Errorf("expected database error, got nil")
		}
	})
}
