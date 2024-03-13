package user

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/google/uuid"
)

type Service interface {
	ListUser() ([]*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(userId uuid.UUID) error
	Login(user *model.User) (*model.User, error)
}
