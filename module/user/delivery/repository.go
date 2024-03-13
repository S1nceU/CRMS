package delivery

import (
	"github.com/S1nceU/CRMS/model"
	"gorm.io/gorm"
)

type UserDelivery struct {
	orm *gorm.DB
}

func NewUserDelivery(orm *gorm.DB) *UserDelivery {
	return &UserDelivery{
		orm: orm,
	}
}

func (u *UserDelivery) ListUser() ([]*model.User, error) {
	var users []*model.User
	err := u.orm.Find(&users).Error
	return users, err
}

func (u *UserDelivery) CreateUser(user *model.User) (*model.User, error) {
	err := u.orm.Create(&user).Error
	return user, err
}

func (u *UserDelivery) UpdateUser(user *model.User) (*model.User, error) {
	err := u.orm.Model(user).Where("user_id = ?", user.UserId).Updates(&user).Error
	return user, err
}

func (u *UserDelivery) DeleteUser(user *model.User) error {
	err := u.orm.Where("user_id = ?", user.UserId).Delete(&user).Error
	return err
}

func (u *UserDelivery) GetUserByUserId(user *model.User) (*model.User, error) {
	err := u.orm.Where("user_id = ?", user.UserId).Find(&user).Error
	return user, err
}

func (u *UserDelivery) GetUserByUsername(user *model.User) (*model.User, error) {
	err := u.orm.Where("username = ?", user.Username).Find(&user).Error
	return user, err
}
