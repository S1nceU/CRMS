package history

import (
	"github.com/S1nceU/CRMS/model"
	"time"
)

type Service interface {
	ListHistories() ([]model.History, error)                   // Get all History
	GetHistoryByID(in int) ([]model.History, error)            // Get History by ID
	GetHistoriesForDate(in time.Time) ([]model.History, error) // Get History by Date
	GetHistoryByHistoryId(in int) (*model.History, error)      // Get History by HistoryId
	CreateHistory(in *model.History) (*model.History, error)   // Create a new History
	UpdateHistory(in *model.History) (*model.History, error)   // Update History data
	DeleteHistory(in int) error                                // Delete History by ID
	DeleteHistoriesByCustomer(in int) error                    // Delete History by CustomerID
	GetHistoryForDuring(in1 string, in2 string) ([]model.History, error)
}
