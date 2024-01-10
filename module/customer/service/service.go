package service

import (
	"errors"
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/customer"
	"github.com/google/uuid"
)

type CustomerService struct {
	repo customer.Repository
}

func NewCustomerService(repo customer.Repository) customer.Service {
	return &CustomerService{
		repo: repo,
	}
}

func (u *CustomerService) ListCustomers() ([]model.Customer, error) {
	var err error
	var customers []*model.Customer
	if customers, err = u.repo.ListCustomers(); err != nil {
		return nil, err
	}
	return convertToSliceOfCustomer(customers), err
}

func (u *CustomerService) ListCustomersByCitizenship(citizenship string) ([]model.Customer, error) {
	var err error
	var customers []*model.Customer
	newCustomer := &model.Customer{
		Citizenship: citizenship,
	}
	if customers, err = u.repo.ListCustomersForCitizenship(newCustomer); err != nil {
		return nil, err
	}
	return convertToSliceOfCustomer(customers), err
}

func (u *CustomerService) ListCustomersByCustomerName(name string) ([]model.Customer, error) {
	var err error
	var customers []*model.Customer

	if name == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}

	newCustomer := &model.Customer{
		Name: name,
	}
	if customers, err = u.repo.ListCustomersByCustomerName(newCustomer); err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, errors.New("error CRMS : There is no this customer")
	}

	return convertToSliceOfCustomer(customers), err
}

func (u *CustomerService) ListCustomersByCustomerPhone(phone string) ([]model.Customer, error) {
	var err error
	var customers []*model.Customer

	if phone == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}

	newCustomer := &model.Customer{
		PhoneNumber: phone,
	}

	if customers, err = u.repo.ListCustomersByCustomerPhone(newCustomer); err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, errors.New("error CRMS : There is no this customer")
	}

	return convertToSliceOfCustomer(customers), err
}

func (u *CustomerService) GetCustomerByID(id string) (*model.Customer, error) {
	var err error
	newCustomer := &model.Customer{
		ID: id,
	}
	if newCustomer, err = u.repo.GetCustomerByID(newCustomer); err != nil {
		return nil, err
	} else if newCustomer.CustomerId == uuid.Nil {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	return newCustomer, err
}

func (u *CustomerService) GetCustomerByCustomerId(customerId uuid.UUID) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer

	newCustomer = &model.Customer{
		CustomerId: customerId,
	}
	if newCustomer, err = u.repo.GetCustomerByCustomerId(newCustomer); err != nil {
		return nil, err
	} else if newCustomer.Name == "" {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	return newCustomer, err
}

func (u *CustomerService) CreateCustomer(customer *model.Customer) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer

	if err = validateCustomerInfo(customer); err != nil {
		return nil, err
	}

	if newCustomer, err = u.repo.GetCustomerByID(customer); err != nil {
		return nil, err
	} else if newCustomer.CustomerId != uuid.Nil {
		return nil, errors.New("error CRMS : This customer is already existed")
	} else {
		customer.CustomerId = uuid.New()
		newCustomer, err = u.repo.CreateCustomer(newCustomer)
		return newCustomer, err
	}
}

func (u *CustomerService) UpdateCustomer(customer *model.Customer) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer

	if _, err = u.GetCustomerByCustomerId(customer.CustomerId); err != nil {
		return nil, err
	}

	if err = validateCustomerInfo(customer); err != nil {
		return nil, err
	}

	if newCustomer, err = u.repo.UpdateCustomer(customer); err != nil {
		return nil, err
	}
	return newCustomer, err
}

func (u *CustomerService) DeleteCustomer(customerId uuid.UUID) error {
	var err error
	newCustomer := &model.Customer{
		CustomerId: customerId,
	}
	if _, err = u.GetCustomerByCustomerId(newCustomer.CustomerId); err != nil {
		return err
	}
	if err = u.repo.DeleteCustomer(newCustomer); err != nil {
		return err
	}
	return nil
}

func convertToSliceOfCustomer(customers []*model.Customer) []model.Customer {
	var customersSlice []model.Customer
	for i := 0; i < len(customers); i++ {
		customersSlice = append(customersSlice, *customers[i])
	}
	return customersSlice
}

func validateCustomerInfo(customer *model.Customer) error {
	if customer.Name == "" {
		return errors.New("error CRMS : Customer Info is incomplete")
	}
	if customer.Gender != "Male" && customer.Gender != "Female" {
		return errors.New("error CRMS : Customer Info is incomplete")
	}
	if customer.Birthday.IsZero() {
		return errors.New("error CRMS : Customer Info is incomplete")
	}
	if customer.ID == "" {
		return errors.New("error CRMS : Customer Info is incomplete")
	}
	if customer.Citizenship == "" {
		return errors.New("error CRMS : Customer Info is incomplete")
	}
	return nil
}
