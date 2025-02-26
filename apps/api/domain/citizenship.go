package domain

import "github.com/S1nceU/CRMS/model"

// CitizenshipRepository is an interface for citizenship repository
type CitizenshipRepository interface {
	ListCitizenships() ([]*model.Citizenship, error)                                            // Get all Citizenships
	GetCitizenshipByID(citizenship *model.Citizenship) (*model.Citizenship, error)              // Get Citizenship by ID
	GetCitizenshipByCitizenshipName(citizenship *model.Citizenship) (*model.Citizenship, error) // Get Citizenship by CitizenshipName
}

// CitizenshipService is an interface for citizenship service
type CitizenshipService interface {
	ListCitizenships() ([]model.Citizenship, error)                                     // Get all Citizenships
	GetCitizenshipByID(id int) (*model.Citizenship, error)                              // Get Citizenship by ID
	GetCitizenshipByCitizenshipName(citizenshipName string) (*model.Citizenship, error) // Get Citizenship by CitizenshipName
}
