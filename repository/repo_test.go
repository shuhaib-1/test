package repository_test

import (
	"math/rand"
	"test/domain"
	"test/repository"
	"testing"
)

func TestRepo(t *testing.T) {

	repo := repository.NewUserRepository()
	t.Run("TestSaveUserAndFindByID", func(t *testing.T) {

		randID := rand.Intn(100000)

		us := domain.User{ID: randID, Name: "shuhaib", Email: "shuhaibpa85@gmail.com"}

		repo.Save(us)

		user, err := repo.FindByID(randID)

		if err != nil {
			t.Fatal("unexpected error:", err)
		}

		if user.Email != "shuhaibpa85@gmail.com" {
			t.Errorf("expected email shuhaibpa85@gmail.com, got %v", user.Email)
		}
	})

	t.Run("TestUserNotFound", func(t *testing.T) {

		_, err := repo.FindByID(1945687)

		if err == nil {
			t.Fatal("expected error is missing, got nil")
		}

		if err.Error() != "user not found" {
			t.Errorf("unexpected error: %v", err)
		}
	})

}
