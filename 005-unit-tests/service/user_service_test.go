package service

import (
	"errors"
	"github.com/prokhorind/go_course/005-unit-tests/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock implements the service.UserRepository interface
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetUserByID(id int) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func TestFetchUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockUser := &models.User{ID: 1, Name: "Test User"}

	mockRepo.On("GetUserByID", 1).Return(mockUser, nil)

	svc := NewUserService(mockRepo)

	user, err := svc.FetchUser(1)

	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)
	mockRepo.AssertExpectations(t)
}

func TestFetchUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockRepo.On("GetUserByID", 999).Return(nil, errors.New("not found"))

	svc := NewUserService(mockRepo)

	user, err := svc.FetchUser(999)

	assert.Nil(t, user)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFetchUser_AdminAccessDenied(t *testing.T) {
	mockRepo := new(MockUserRepo)
	adminUser := &models.User{ID: 5, Name: "AdminJohn"}

	mockRepo.On("GetUserByID", 5).Return(adminUser, nil)

	svc := NewUserService(mockRepo)

	user, err := svc.FetchUser(5)

	assert.Nil(t, user)
	assert.Error(t, err)
	assert.EqualError(t, err, "access to admin users is restricted")
	mockRepo.AssertExpectations(t)
}
