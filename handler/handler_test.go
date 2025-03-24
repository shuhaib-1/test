package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"test/domain"
	"test/handler"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUsecase) GetUserById(id int) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestCreateUserHandler(t *testing.T) {
	mockUsecase := new(MockUsecase)
	handler := handler.NewUserHandler(mockUsecase)

	user := domain.User{Name: "shuhaib", Email: "shuhaib@gmail.com"}
	mockUsecase.On("CreateUser", user).Return(nil)

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
	mockUsecase.AssertCalled(t, "CreateUser", user)
}
