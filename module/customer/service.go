package customer

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/google/uuid"
)

type Service interface {
	ListCustomers() ([]model.Customer, error)                      // Get all Customers
	GetCustomersByCitizenship(in string) ([]model.Customer, error) // Get all Customers by citizenship
	GetCustomerByID(in string) (*model.Customer, error)            // Get Customer by ID
	GetCustomerByCustomerId(in uuid.UUID) (*model.Customer, error) // Get Customer by customer_id
	CreateCustomer(in *model.Customer) (*model.Customer, error)    // Create a new Customer
	UpdateCustomer(in *model.Customer) (*model.Customer, error)    // Update customer data
	DeleteCustomer(in uuid.UUID) error                             // Delete Customer by customer_id
	GetCustomerByCustomerName(in string) ([]model.Customer, error) // Get Customer by customer_name
}
