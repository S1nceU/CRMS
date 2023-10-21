package repository

import (
	"crms/model"
	"crms/module/customer"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	orm *gorm.DB
}

func NewCustomerRepository(orm *gorm.DB) customer.Repository {
	return &CustomerRepository{
		orm: orm,
	}
}

func (u *CustomerRepository) GetCustomerList() ([]*model.Customer, error) {
	var err error
	var in []*model.Customer
	err = u.orm.Find(&in).Error
	return in, err
}

func (u *CustomerRepository) GetCustomerListForCitizenship(in *model.Customer) ([]*model.Customer, error) {
	var err error
	var out []*model.Customer
	if err = u.orm.Where("Citizenship = ?", in.Citizenship).Find(&out).Error; err != nil {
		return nil, err
	}
	return out, err
}

func (u *CustomerRepository) GetCustomer(in *model.Customer) (*model.Customer, error) {
	var err error
	err = u.orm.Where("ID = ?", in.ID).Find(&in).Error
	return in, err
}

func (u *CustomerRepository) GetCustomerForCID(in *model.Customer) (*model.Customer, error) {
	var err error
	err = u.orm.Where("customer_id = ?", in.CustomerId).Find(&in).Error
	return in, err
}

func (u *CustomerRepository) CreateCustomer(in *model.Customer) (*model.Customer, error) {
	var err error
	err = u.orm.Create(in).Error
	return in, err
}

func (u *CustomerRepository) UpdateCustomer(in *model.Customer) (*model.Customer, error) {
	var err error
	err = u.orm.Model(in).Where("customer_id = ?", in.CustomerId).Updates(&in).Error
	return in, err
}

func (u *CustomerRepository) DeleteCustomer(in *model.Customer) error {
	var err error
	err = u.orm.Where("customer_id = ?", in.CustomerId).Delete(&in).Error
	return err
}
