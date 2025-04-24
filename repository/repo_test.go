package repository_test

import (
	"test/domain"
	"test/repository"
	"testing"
)

func TestRepo(t *testing.T) {

	repo := repository.NewUserRepository()
	t.Run("TestSaveUserAndFindByID", func(t *testing.T) {

		us := domain.User{ID: 999, Name: "shuhaib", Email: "shuhaibpa85@gmail.com"}

		repo.Save(us)

		user, err := repo.FindByID(999)

		if err != nil {
			t.Fatal("unexpected error:", err)
		}

		if user.Email != "shuhaibpa85@gmail.com" {
			t.Errorf("expected email shuhaibpa85@gmail.com, got %v", user.Email)
		}
	})

	t.Run("TestUserNotFound", func(t *testing.T) {

		_, err := repo.FindByID(19456)

		if err == nil {
			t.Fatal("expected error is missing, got nil")
		}

		if err.Error() != "user not found" {
			t.Errorf("unexpected error: %v", err)
		}
	})

}
