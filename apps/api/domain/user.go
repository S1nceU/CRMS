package domain

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/google/uuid"
)

// UserRepository is an interface for user repository
type UserRepository interface {
	ListUser() ([]*model.User, error)                        // Get all User
	CreateUser(user *model.User) (*model.User, error)        // Create a new User
	UpdateUser(user *model.User) (*model.User, error)        // Update User data
	DeleteUser(user *model.User) error                       // Delete User by UserID
	GetUserByUserId(user *model.User) (*model.User, error)   // Get User by UserID
	GetUserByUsername(user *model.User) (*model.User, error) // Get User by Username
}

// UserService is an interface for user service
type UserService interface {
	ListUser() ([]*model.User, error)                       // Get all User
	CreateUser(user *model.User) (*model.User, error)       // Create a new User
	UpdateUser(user *model.User) (*model.User, error)       // Update User data
	DeleteUser(userId uuid.UUID) error                      // Delete User by UserID
	GetUserByUserId(userId uuid.UUID) (*model.User, error)  // Get User by UserID
	GetUserByUsername(username string) (*model.User, error) // Get User by Username
	Login(username, password string) (string, error)        // Login User
	Authentication(tokenString string) (string, error)      // Authentication user Token
}
