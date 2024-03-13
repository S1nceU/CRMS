package user

import "github.com/S1nceU/CRMS/model"

type Repository interface {
	ListUser() ([]*model.User, error)                        // Get all User
	CreateUser(user *model.User) (*model.User, error)        // Create a new User
	UpdateUser(user *model.User) (*model.User, error)        // Update User data
	DeleteUser(user *model.User) error                       // Delete User by UserID
	GetUserByUserId(user *model.User) (*model.User, error)   // Get User by UserID
	GetUserByUsername(user *model.User) (*model.User, error) // Get User by Username
}
