package service

import (
	"strings"
	"testing"

	"github.com/S1nceU/CRMS/apps/api/config"
	"github.com/S1nceU/CRMS/apps/api/model"
	"github.com/google/uuid"
)

type fakeUserRepository struct {
	usersByUsername map[string]*model.User
}

func (f *fakeUserRepository) ListUser() ([]*model.User, error) {
	return nil, nil
}

func (f *fakeUserRepository) CreateUser(user *model.User) (*model.User, error) {
	return user, nil
}

func (f *fakeUserRepository) UpdateUser(user *model.User) (*model.User, error) {
	return user, nil
}

func (f *fakeUserRepository) DeleteUser(user *model.User) error {
	return nil
}

func (f *fakeUserRepository) GetUserByUserId(user *model.User) (*model.User, error) {
	for _, storedUser := range f.usersByUsername {
		if storedUser.Id == user.Id {
			return storedUser, nil
		}
	}
	return nil, nil
}

func (f *fakeUserRepository) GetUserByUsername(user *model.User) (*model.User, error) {
	storedUser, ok := f.usersByUsername[user.Username]
	if !ok {
		return nil, nil
	}
	return storedUser, nil
}

func TestLoginReturnsTokenThatAuthenticatesUsername(t *testing.T) {
	originalConfig := config.Val
	t.Cleanup(func() {
		config.Val = originalConfig
	})

	config.Val = config.Config{
		TokenSecret:      "test-secret",
		TokenIssuer:      "crms-test",
		TokenExpireHours: 1,
	}

	repo := &fakeUserRepository{
		usersByUsername: map[string]*model.User{
			"alice": {
				Id:       uuid.New(),
				Username: "alice",
				Password: "correct-password",
			},
		},
	}
	userService := NewUserService(repo)

	token, err := userService.Login("alice", "correct-password")

	if err != nil {
		t.Fatalf("Login() error = %v", err)
	}
	if strings.Count(token, ".") != 2 {
		t.Fatalf("Login() token = %q, want a JWT with three segments", token)
	}

	username, err := userService.Authentication(token)
	if err != nil {
		t.Fatalf("Authentication() error = %v", err)
	}
	if username != "alice" {
		t.Fatalf("Authentication() username = %q, want %q", username, "alice")
	}
}

func TestLoginRejectsIncorrectPassword(t *testing.T) {
	originalConfig := config.Val
	t.Cleanup(func() {
		config.Val = originalConfig
	})

	config.Val = config.Config{
		TokenSecret:      "test-secret",
		TokenIssuer:      "crms-test",
		TokenExpireHours: 1,
	}

	repo := &fakeUserRepository{
		usersByUsername: map[string]*model.User{
			"alice": {
				Id:       uuid.New(),
				Username: "alice",
				Password: "correct-password",
			},
		},
	}
	userService := NewUserService(repo)

	token, err := userService.Login("alice", "wrong-password")

	if token != "" {
		t.Fatalf("Login() token = %q, want empty token", token)
	}
	if err == nil {
		t.Fatal("Login() error = nil, want password is incorrect")
	}
	if err.Error() != "password is incorrect" {
		t.Fatalf("Login() error = %v, want password is incorrect", err)
	}
}
