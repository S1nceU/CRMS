package domain

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/google/uuid"
)

// CustomerRepository is an interface for customer repository
type CustomerRepository interface {
	ListCustomers() ([]*model.Customer, error)                                      // Get all Customers
	ListCustomersByCitizenship(customer *model.Customer) ([]*model.Customer, error) // Get all Customers by Citizenship
	ListCustomersByName(customer *model.Customer) ([]*model.Customer, error)        // Get Customer by CustomerName
	ListCustomersByPhone(customer *model.Customer) ([]*model.Customer, error)       // Get Customer by CustomerPhone
	GetCustomerByNationalId(customer *model.Customer) (*model.Customer, error)      // Get Customer by ID
	GetCustomerByCustomerId(customer *model.Customer) (*model.Customer, error)      // Get Customer by CustomerId
	CreateCustomer(customer *model.Customer) (*model.Customer, error)               // Create a new Customer
	UpdateCustomer(customer *model.Customer) (*model.Customer, error)               // Update Customer data
	DeleteCustomer(customer *model.Customer) error                                  // Delete Customer by CustomerId
}

// CustomerService is an interface for customer service
type CustomerService interface {
	ListCustomers() ([]*model.Customer, error)                             // Get all Customers
	ListCustomersByCitizenship(citizenship int) ([]*model.Customer, error) // Get all Customers by citizenship
	ListCustomersByName(name string) ([]*model.Customer, error)            // Get Customer by customer_name
	ListCustomersByPhone(phone string) ([]*model.Customer, error)          // Get Customer by customer_phone
	GetCustomerByNationalId(id string) (*model.Customer, error)            // Get Customer by ID
	GetCustomerByCustomerId(customerId uuid.UUID) (*model.Customer, error) // Get Customer by customer_id
	CreateCustomer(customer *model.Customer) (*model.Customer, error)      // Create a new Customer
	UpdateCustomer(customer *model.Customer) (*model.Customer, error)      // Update customer data
	DeleteCustomer(customerId uuid.UUID) error                             // Delete Customer by customer_id
}
