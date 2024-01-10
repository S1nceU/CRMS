package history

import "github.com/S1nceU/CRMS/model"

type Repository interface {
	ListHistories() ([]*model.History, error)                                                          // Get all History
	ListHistoriesByCustomer(history *model.History) ([]*model.History, error)                          // Get History by CustomerId
	ListHistoriesForDate(history *model.History) ([]*model.History, error)                             // Get History by Date
	ListHistoriesForDuring(history1 *model.History, history2 *model.History) ([]*model.History, error) // Get History by During
	GetHistoryByHistoryId(history *model.History) (*model.History, error)                              // Get History by HistoryID
	CreateHistory(history *model.History) (*model.History, error)                                      // Create a new History
	UpdateHistory(history *model.History) (*model.History, error)                                      // Update History data
	DeleteHistory(history *model.History) error                                                        // Delete History by HistoryID
	DeleteHistoriesByCustomer(history *model.History) error                                            // Delete History by CustomerID
	ConfirmCustomerExistence(customer *model.Customer) (*model.Customer, error)                        // Confirm Customer Existed
}
