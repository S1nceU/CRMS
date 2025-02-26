package service

import (
	"errors"
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
	"github.com/google/uuid"
)

type CustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) domain.CustomerService {
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

func (u *CustomerService) ListCustomersByCitizenship(citizenship int) ([]model.Customer, error) {
	var err error
	var customers []*model.Customer
	newCustomer := &model.Customer{
		CitizenshipId: citizenship,
	}
	if customers, err = u.repo.ListCustomersByCitizenship(newCustomer); err != nil {
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

func (u *CustomerService) GetCustomerByNationalId(id string) (*model.Customer, error) {
	var err error
	newCustomer := &model.Customer{
		NationalId: id,
	}
	if newCustomer, err = u.repo.GetCustomerByNationalId(newCustomer); err != nil {
		return nil, err
	} else if newCustomer.Id == uuid.Nil {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	return newCustomer, err
}

func (u *CustomerService) GetCustomerByCustomerId(customerId uuid.UUID) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer

	newCustomer = &model.Customer{
		Id: customerId,
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
	if newCustomer, err = u.repo.GetCustomerByNationalId(customer); err != nil {
		return nil, err
	} else if newCustomer.Id != uuid.Nil {
		return nil, errors.New("error CRMS : This customer is already existed")
	} else {
		customer.Id = uuid.New()
		newCustomer, err = u.repo.CreateCustomer(newCustomer)
		return newCustomer, err
	}
}

func (u *CustomerService) UpdateCustomer(customer *model.Customer) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer

	if _, err = u.GetCustomerByCustomerId(customer.Id); err != nil {
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
		Id: customerId,
	}
	if _, err = u.GetCustomerByCustomerId(newCustomer.Id); err != nil {
		return err
	}
	if err = u.repo.DeleteCustomer(newCustomer); err != nil {
		return err
	}
	return nil
}

func convertToSliceOfCustomer(customers []*model.Customer) []model.Customer {
	var customersSlice []model.Customer
	for _, customer := range customers {
		customersSlice = append(customersSlice, *customer)
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
	if customer.NationalId == "" {
		return errors.New("error CRMS : Customer Info is incomplete")
	}
	if customer.CitizenshipId == 0 {
		return errors.New("error CRMS : Customer Info is incomplete")
	}
	return nil
}
