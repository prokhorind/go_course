package service

import (
	"errors"
	"github.com/prokhorind/go_course/005-unit-tests/models"
	"testing"
)

// ---- Manual mock: implements UserRepository interface ----

type fakeUserRepo struct {
	userToReturn *models.User
	errToReturn  error
}

func (f *fakeUserRepo) GetUserByID(id int) (*models.User, error) {
	return f.userToReturn, f.errToReturn
}

// ---- Unit Tests ----

func TestFetchUser_WithFakeRepo_Success(t *testing.T) {
	fake := &fakeUserRepo{
		userToReturn: &models.User{ID: 1, Name: "Manual Fake"},
		errToReturn:  nil,
	}

	svc := NewUserService(fake)

	user, err := svc.FetchUser(1)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if user == nil || user.ID != 1 || user.Name != "Manual Fake" {
		t.Errorf("Unexpected user: %+v", user)
	}
}

func TestFetchUser_WithFakeRepo_Error(t *testing.T) {
	fake := &fakeUserRepo{
		userToReturn: nil,
		errToReturn:  errors.New("user not found"),
	}

	svc := NewUserService(fake)

	user, err := svc.FetchUser(999)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if user != nil {
		t.Errorf("Expected nil user, got: %+v", user)
	}
}
