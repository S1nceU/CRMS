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

func (u *CustomerService) GetCustomerList() ([]model.Customer, error) {
	var err error
	var point []*model.Customer
	var out []model.Customer
	if point, err = u.repo.GetCustomerList(); err != nil {
		return nil, err
	}
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *CustomerService) GetCustomerListForCitizenship(in string) ([]model.Customer, error) {
	var err error
	var point []*model.Customer
	var out []model.Customer
	newCustomer := &model.Customer{
		Citizenship: in,
	}
	if point, err = u.repo.GetCustomerListForCitizenship(newCustomer); err != nil {
		return nil, err
	}
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *CustomerService) GetCustomer(in string) (*model.Customer, error) {
	var err error
	newCustomer := &model.Customer{
		ID: in,
	}
	if newCustomer, err = u.repo.GetCustomer(newCustomer); err != nil {
		return nil, err
	} else if newCustomer.CustomerId == 0 {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	return newCustomer, err
}

func (u *CustomerService) GetCustomerForCID(in int) (*model.Customer, error) {
	var err error
	newCustomer := &model.Customer{
		CustomerId: in,
	}
	if newCustomer, err = u.repo.GetCustomerForCID(newCustomer); err != nil {
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
	if in.Birthday == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.ID == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Citizenship == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}

	if newCustomer, err = u.repo.GetCustomer(in); err != nil {
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
	if in.Name == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Gender != "Male" && in.Gender != "Female" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Birthday == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.ID == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}
	if in.Citizenship == "" {
		return nil, errors.New("error CRMS : Customer Info is incomplete")
	}

	if _, err = u.GetCustomerForCID(in.CustomerId); err != nil {
		return nil, err
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
	if _, err = u.GetCustomerForCID(newCustomer.CustomerId); err != nil {
		return err
	}
	if err = u.repo.DeleteCustomer(newCustomer); err != nil {
		return err
	}
	return nil
}
