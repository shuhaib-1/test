package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"test/domain"
	"test/handler"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock implementation of UserUseCase
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
	app := fiber.New()

	mockUsecase := new(MockUsecase)
	h := handler.NewUserHandler(mockUsecase)
	app.Post("/user", h.CreateUser)

	user := domain.User{Name: "shuhaib", Email: "shuhaib@gmail.com"}
	mockUsecase.On("CreateUser", user).Return(nil)

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusCreated, resp.StatusCode)

	mockUsecase.AssertCalled(t, "CreateUser", user)
}

func TestCreateUserHandler_InvalidJSON(t *testing.T) {
	app := fiber.New()

	mockUsecase := new(MockUsecase)
	h := handler.NewUserHandler(mockUsecase)
	app.Post("/user", h.CreateUser)

	body := `{"name": "shuhaib", "email":}` // invalid JSON
	req := httptest.NewRequest("POST", "/user", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestCreateUserHandler_MissingFields(t *testing.T) {
	app := fiber.New()

	mockUsecase := new(MockUsecase)
	h := handler.NewUserHandler(mockUsecase)
	app.Post("/user", h.CreateUser)

	user := domain.User{Name: "", Email: ""}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
