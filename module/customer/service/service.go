package service

import (
	"errors"
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/customer"
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
	var point []*model.Customer
	var out []model.Customer
	if point, err = u.repo.ListCustomers(); err != nil {
		return nil, err
	}
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *CustomerService) GetCustomersByCitizenship(in string) ([]model.Customer, error) {
	var err error
	var point []*model.Customer
	var out []model.Customer
	newCustomer := &model.Customer{
		Citizenship: in,
	}
	if point, err = u.repo.ListCustomersForCitizenship(newCustomer); err != nil {
		return nil, err
	}
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *CustomerService) GetCustomerByID(in string) (*model.Customer, error) {
	var err error
	newCustomer := &model.Customer{
		ID: in,
	}
	if newCustomer, err = u.repo.GetCustomerByID(newCustomer); err != nil {
		return nil, err
	} else if newCustomer.CustomerId == 0 {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	return newCustomer, err
}

func (u *CustomerService) GetCustomerByCustomerId(in int) (*model.Customer, error) {
	var err error
	newCustomer := &model.Customer{
		CustomerId: in,
	}
	if newCustomer, err = u.repo.GetCustomerByCustomerId(newCustomer); err != nil {
		return nil, err
	} else if newCustomer.Name == "" {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	return newCustomer, err
}

func (u *CustomerService) CreateCustomer(in *model.Customer) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer

	if in.Name == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Gender != "Male" && in.Gender != "Female" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Birthday.IsZero() {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.ID == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Citizenship == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}

	if newCustomer, err = u.repo.GetCustomerByID(in); err != nil {
		return nil, err
	} else if newCustomer.CustomerId != 0 {
		return nil, errors.New("error CRMS : This customer is already existed")
	} else {
		newCustomer, err = u.repo.CreateCustomer(newCustomer)
		return newCustomer, err
	}
}

func (u *CustomerService) UpdateCustomer(in *model.Customer) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer

	if _, err = u.GetCustomerByCustomerId(in.CustomerId); err != nil {
		return nil, err
	}

	if in.Name == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Gender != "Male" && in.Gender != "Female" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Birthday.IsZero() {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.ID == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Citizenship == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}

	if newCustomer, err = u.repo.UpdateCustomer(in); err != nil {
		return nil, err
	}
	return newCustomer, err
}

func (u *CustomerService) DeleteCustomer(in int) error {
	var err error
	newCustomer := &model.Customer{
		CustomerId: in,
	}
	if _, err = u.GetCustomerByCustomerId(newCustomer.CustomerId); err != nil {
		return err
	}
	if err = u.repo.DeleteCustomer(newCustomer); err != nil {
		return err
	}
	return nil
}
