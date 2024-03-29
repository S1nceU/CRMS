package customer

import "crms/model"

type Service interface {
	GetCustomerList() ([]model.Customer, error)                        // Get all Customer
	GetCustomerListForCitizenship(in string) ([]model.Customer, error) // Get all Customer by citizenship
	GetCustomer(in string) (*model.Customer, error)                    // Get Customer by ID and citizenship
	GetCustomerForCID(in int) (*model.Customer, error)                 // Get Customer by customer_id
	CreateCustomer(in []byte) (*model.Customer, error)                 // Create a new customer
	//UpdateCustomer(in *model.Customer) (*model.Customer, error)         // Update customer data by a whole customer
	//DeleteCustomer(in int) error                                        // Delete customer by customer_id
}
