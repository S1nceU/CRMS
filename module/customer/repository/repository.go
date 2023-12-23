package repository

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/customer"
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

func (u *CustomerRepository) ListCustomers() ([]*model.Customer, error) {
	var err error
	var in []*model.Customer
	err = u.orm.Find(&in).Error
	return in, err
}

func (u *CustomerRepository) ListCustomersForCitizenship(in *model.Customer) ([]*model.Customer, error) {
	var err error
	var out []*model.Customer
	if err = u.orm.Where("Citizenship = ?", in.Citizenship).Find(&out).Error; err != nil {
		return nil, err
	}
	return out, err
}

func (u *CustomerRepository) GetCustomerByID(in *model.Customer) (*model.Customer, error) {
	var err error
	err = u.orm.Where("ID = ?", in.ID).Find(&in).Error
	return in, err
}

func (u *CustomerRepository) GetCustomerByCustomerId(in *model.Customer) (*model.Customer, error) {
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

func (u *CustomerRepository) GetCustomerByCustomerName(in *model.Customer) ([]*model.Customer, error) {
	var err error
	var out []*model.Customer
	err = u.orm.Where("Name LIKE ?", "%"+in.Name+"%").Find(&out).Error
	return out, err
}
