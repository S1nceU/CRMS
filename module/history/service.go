package history

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/google/uuid"
)

type Service interface {
	ListHistories() ([]model.History, error)                             // Get all History
	GetHistoryByCustomerId(in uuid.UUID) ([]model.History, error)        // Get History by CustomerId
	GetHistoriesForDate(in string) ([]model.History, error)              // Get History by Date
	GetHistoryByHistoryId(in uuid.UUID) (*model.History, error)          // Get History by HistoryId
	CreateHistory(in *model.History) (*model.History, error)             // Create a new History
	UpdateHistory(in *model.History) (*model.History, error)             // Update History data
	DeleteHistory(in uuid.UUID) error                                    // Delete History by ID
	DeleteHistoriesByCustomer(in uuid.UUID) error                        // Delete History by CustomerID
	GetHistoryForDuring(in1 string, in2 string) ([]model.History, error) // Get History by During
}
