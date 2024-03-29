package service

import (
	"crms/model"
	"crms/module/customer"
	"encoding/json"
	"errors"
)

type CustomerService struct {
	repo customer.Repository
}

func NewCustomer(repo customer.Repository) customer.Service {
	return &CustomerService{
		repo: repo,
	}
}

func (u *CustomerService) GetCustomerList() ([]model.Customer, error) {
	var err error
	var point []*model.Customer
	var out []model.Customer
	point, err = u.repo.GetCustomerList()
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
	point, err = u.repo.GetCustomerListForCitizenship(newCustomer)
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
	newCustomer, err = u.repo.GetCustomer(newCustomer)
	return newCustomer, err
}

func (u *CustomerService) GetCustomerForCID(in int) (*model.Customer, error) {
	var err error
	newCustomer := &model.Customer{
		Customer_id: in,
	}
	newCustomer, err = u.repo.GetCustomerForCID(newCustomer)
	return newCustomer, err
}

func (u *CustomerService) CreateCustomer(in []byte) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer
	json.Unmarshal(in, &newCustomer)
	if newCustomer, err = u.repo.GetCustomer(newCustomer); err != nil {
		newCustomer, err = u.repo.CreateCustomer(newCustomer)
		return newCustomer, err
	} else {
		return nil, errors.New("Error CRMS : This customer is already existed.")
	}
}
