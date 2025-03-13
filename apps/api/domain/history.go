package domain

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/google/uuid"
)

// HistoryRepository is an interface for History repository
type HistoryRepository interface {
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

// HistoryService is an interface for History service
type HistoryService interface {
	ListHistories() ([]*model.History, error)                                // Get all History
	ListHistoriesByCustomerId(in uuid.UUID) ([]*model.History, error)        // Get History by CustomerId
	ListHistoriesForDate(in string) ([]*model.History, error)                // Get History by Date
	ListHistoriesForDuring(in1 string, in2 string) ([]*model.History, error) // Get History by During
	GetHistoryByHistoryId(in uuid.UUID) (*model.History, error)              // Get History by HistoryId
	CreateHistory(in *model.History) (*model.History, error)                 // Create a new History
	UpdateHistory(in *model.History) (*model.History, error)                 // Update History data
	DeleteHistory(in uuid.UUID) error                                        // Delete History by ID
	DeleteHistoriesByCustomer(in uuid.UUID) error                            // Delete History by CustomerID
}
