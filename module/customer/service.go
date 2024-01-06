package customer

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/google/uuid"
)

type Service interface {
	ListCustomers() ([]model.Customer, error)                                // Get all Customers
	ListCustomersByCitizenship(citizenship string) ([]model.Customer, error) // Get all Customers by citizenship
	ListCustomersByCustomerName(name string) ([]model.Customer, error)       // Get Customer by customer_name
	ListCustomersByCustomerPhone(phone string) ([]model.Customer, error)     // Get Customer by customer_phone
	GetCustomerByID(id string) (*model.Customer, error)                      // Get Customer by ID
	GetCustomerByCustomerId(customerId uuid.UUID) (*model.Customer, error)   // Get Customer by customer_id
	CreateCustomer(customer *model.Customer) (*model.Customer, error)        // Create a new Customer
	UpdateCustomer(customer *model.Customer) (*model.Customer, error)        // Update customer data
	DeleteCustomer(customerId uuid.UUID) error                               // Delete Customer by customer_id
}
