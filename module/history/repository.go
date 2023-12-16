package history

import "github.com/S1nceU/CRMS/model"

type Repository interface {
	GetAllHistories() ([]*model.History, error)                                           // Get all History
	GetHistoriesByCustomer(in *model.History) ([]*model.History, error)                   // Get History by CustomerId
	GetHistoriesForDate(in *model.History) ([]*model.History, error)                      // Get History by Date
	GetHistoryByHistoryId(in *model.History) (*model.History, error)                      // Get History by HistoryID
	CreateHistory(in *model.History) (*model.History, error)                              // Create a new History
	UpdateHistory(in *model.History) (*model.History, error)                              // Update History data
	DeleteHistory(in *model.History) error                                                // Delete History by HistoryID
	ConfirmCustomerExistence(in *model.Customer) (*model.Customer, error)                 // Confirm Customer Existed
	DeleteHistoriesByCustomer(in *model.History) error                                    // Delete History by CustomerID
	GetHistoryForDuring(in1 *model.History, in2 *model.History) ([]*model.History, error) // Get History by During
}
