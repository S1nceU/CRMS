package service

import (
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
	"github.com/google/uuid"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
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
		Id: userId,
	}
	err = u.repo.DeleteUser(newUser)
	return err
}

func (u *UserService) GetUserByUserId(userId uuid.UUID) (*model.User, error) {
	var err error
	newUser := &model.User{
		Id: userId,
	}
	newUser, err = u.repo.GetUserByUserId(newUser)
	return newUser, err
}

func (u *UserService) GetUserByUsername(username string) (*model.User, error) {
	var err error
	newUser := &model.User{
		Username: username,
	}
	newUser, err = u.repo.GetUserByUsername(newUser)
	return newUser, err
}
