package repository

import (
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	orm *gorm.DB
}

func NewCustomerRepository(orm *gorm.DB) domain.CustomerRepository {
	return &CustomerRepository{
		orm: orm,
	}
}

func (u *CustomerRepository) ListCustomers() ([]*model.Customer, error) {
	var customers []*model.Customer
	err := u.orm.Find(&customers).Error
	return customers, err
}

func (u *CustomerRepository) ListCustomersByCitizenship(customer *model.Customer) ([]*model.Customer, error) {
	var customers []*model.Customer
	err := u.orm.Where("CitizenshipId = ?", customer.CitizenshipId).Find(&customers).Error
	return customers, err
}

func (u *CustomerRepository) ListCustomersByCustomerName(customer *model.Customer) ([]*model.Customer, error) {
	var customers []*model.Customer
	err := u.orm.Where("Name LIKE ?", "%"+customer.Name+"%").Find(&customers).Error
	return customers, err
}

func (u *CustomerRepository) ListCustomersByCustomerPhone(customer *model.Customer) ([]*model.Customer, error) {
	var customers []*model.Customer
	err := u.orm.Where("PhoneNumber LIKE ?", "%"+customer.PhoneNumber+"%").Find(&customers).Error
	return customers, err
}

func (u *CustomerRepository) GetCustomerByNationalId(customer *model.Customer) (*model.Customer, error) {
	err := u.orm.Where("NationalId = ?", customer.NationalId).Find(&customer).Error
	return customer, err
}

func (u *CustomerRepository) GetCustomerByCustomerId(customer *model.Customer) (*model.Customer, error) {
	err := u.orm.Where("Id = ?", customer.Id).Find(&customer).Error
	return customer, err
}

func (u *CustomerRepository) CreateCustomer(customer *model.Customer) (*model.Customer, error) {
	err := u.orm.Create(customer).Error
	return customer, err
}

func (u *CustomerRepository) UpdateCustomer(customer *model.Customer) (*model.Customer, error) {
	err := u.orm.Model(customer).Where("Id = ?", customer.Id).Updates(&customer).Error
	return customer, err
}

func (u *CustomerRepository) DeleteCustomer(customer *model.Customer) error {
	return u.orm.Where("Id = ?", customer.Id).Delete(&customer).Error
}
