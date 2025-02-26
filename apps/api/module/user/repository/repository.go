package repository

import (
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	orm *gorm.DB
}

func NewUserRepository(orm *gorm.DB) domain.UserRepository {
	return &UserRepository{
		orm: orm,
	}
}

func (u *UserRepository) ListUser() ([]*model.User, error) {
	var users []*model.User
	err := u.orm.Find(&users).Error
	return users, err
}

func (u *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	err := u.orm.Create(&user).Error
	return user, err
}

func (u *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := u.orm.Model(user).Where("user_id = ?", user.Id).Updates(&user).Error
	return user, err
}

func (u *UserRepository) DeleteUser(user *model.User) error {
	err := u.orm.Where("user_id = ?", user.Id).Delete(&user).Error
	return err
}

func (u *UserRepository) GetUserByUserId(user *model.User) (*model.User, error) {
	err := u.orm.Where("user_id = ?", user.Id).Find(&user).Error
	return user, err
}

func (u *UserRepository) GetUserByUsername(user *model.User) (*model.User, error) {
	err := u.orm.Where("username = ?", user.Username).Find(&user).Error
	return user, err
}
