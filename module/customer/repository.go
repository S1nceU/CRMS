package customer

import "github.com/S1nceU/CRMS/model"

type Repository interface {
	ListCustomers() ([]*model.Customer, error)                                 // Get all Customers
	ListCustomersForCitizenship(in *model.Customer) ([]*model.Customer, error) // Get all Customers by Citizenship
	GetCustomerByID(in *model.Customer) (*model.Customer, error)               // Get Customer by ID
	GetCustomerByCustomerId(in *model.Customer) (*model.Customer, error)       // Get Customer by CustomerId
	CreateCustomer(in *model.Customer) (*model.Customer, error)                // Create a new Customer
	UpdateCustomer(in *model.Customer) (*model.Customer, error)                // Update Customer data
	DeleteCustomer(in *model.Customer) error                                   // Delete Customer by CustomerId
	GetCustomerByCustomerName(in *model.Customer) ([]*model.Customer, error)   // Get Customer by CustomerName
}
