package repository

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/user"
	"github.com/google/uuid"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) user.Service {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) ListUser() ([]*model.User, error) {
	return u.repo.ListUser()
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	return u.repo.CreateUser(user)
}

func (u *UserService) UpdateUser(user *model.User) (*model.User, error) {
	return u.repo.UpdateUser(user)
}

func (u *UserService) DeleteUser(userId uuid.UUID) error {
	var err error
	newUser := &model.User{
		UserId: userId,
	}
	err = u.repo.DeleteUser(newUser)
	return err
}

func (u *UserService) Login(user *model.User) (*model.User, error) {
	return u.repo.GetUserByUsername(user)
}
