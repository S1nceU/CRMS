package customer

import "github.com/S1nceU/CRMS/model"

type Repository interface {
	ListCustomers() ([]*model.Customer, error)                                        // Get all Customers
	ListCustomersForCitizenship(customer *model.Customer) ([]*model.Customer, error)  // Get all Customers by Citizenship
	ListCustomersByCustomerName(customer *model.Customer) ([]*model.Customer, error)  // Get Customer by CustomerName
	ListCustomersByCustomerPhone(customer *model.Customer) ([]*model.Customer, error) // Get Customer by CustomerPhone
	GetCustomerByID(customer *model.Customer) (*model.Customer, error)                // Get Customer by ID
	GetCustomerByCustomerId(customer *model.Customer) (*model.Customer, error)        // Get Customer by CustomerId
	CreateCustomer(customer *model.Customer) (*model.Customer, error)                 // Create a new Customer
	UpdateCustomer(customer *model.Customer) (*model.Customer, error)                 // Update Customer data
	DeleteCustomer(customer *model.Customer) error                                    // Delete Customer by CustomerId
}
