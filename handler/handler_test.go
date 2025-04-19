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
	tests := []struct {
		name           string
		requestBody    []byte
		expectedStatus int
		mockSetup      func(mock *MockUsecase)
	}{
		{
			name: "Valid User",
			requestBody: func() []byte {
				user := domain.User{Name: "shuhaib", Email: "shuhaib@gmail.com"}
				body, _ := json.Marshal(user)
				return body
			}(),
			expectedStatus: http.StatusCreated,
			mockSetup: func(mock *MockUsecase) {
				mock.On("CreateUser", domain.User{Name: "shuhaib", Email: "shuhaib@gmail.com"}).Return(nil)
			},
		},
		{
			name:           "Invalid JSON",
			requestBody:    []byte(`{"name": "shuhaib", "email":}`),
			expectedStatus: http.StatusBadRequest,
			mockSetup:      func(mock *MockUsecase) {}, // No usecase method expected
		},
		{
			name: "Missing Fields",
			requestBody: func() []byte {
				user := domain.User{Name: "", Email: ""}
				body, _ := json.Marshal(user)
				return body
			}(),
			expectedStatus: http.StatusBadRequest,
			mockSetup:      func(mock *MockUsecase) {}, // No usecase method expected
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := new(MockUsecase)
			tc.mockSetup(mockUsecase)

			handler := handler.NewUserHandler(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(tc.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handler.CreateUser(rec, req)

			require.Equal(t, tc.expectedStatus, rec.Code)

			if tc.name == "Valid User" {
				mockUsecase.AssertCalled(t, "CreateUser", domain.User{Name: "shuhaib", Email: "shuhaib@gmail.com"})
			}
		})
	}
}
