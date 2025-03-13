package service

import (
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
)

type CitizenshipService struct {
	repo domain.CitizenshipRepository
}

func NewCitizenshipService(repo domain.CitizenshipRepository) domain.CitizenshipService {
	return &CitizenshipService{
		repo: repo,
	}
}

func (u *CitizenshipService) ListCitizenships() ([]*model.Citizenship, error) {
	var err error
	var citizenships []*model.Citizenship
	if citizenships, err = u.repo.ListCitizenships(); err != nil {
		return nil, err
	}
	return convertToSliceOfCitizenship(citizenships), err
}

func (u *CitizenshipService) GetCitizenshipByID(id int) (*model.Citizenship, error) {
	var err error
	var citizenship *model.Citizenship
	newCitizenship := &model.Citizenship{
		Id: id,
	}
	if citizenship, err = u.repo.GetCitizenshipByID(newCitizenship); err != nil {
		return nil, err
	}
	return citizenship, err
}

func (u *CitizenshipService) GetCitizenshipByCitizenshipName(citizenshipName string) (*model.Citizenship, error) {
	var err error
	var citizenship *model.Citizenship
	newCitizenship := &model.Citizenship{
		Nation: citizenshipName,
	}
	if citizenship, err = u.repo.GetCitizenshipByCitizenshipName(newCitizenship); err != nil {
		return nil, err
	}
	return citizenship, err
}
func convertToSliceOfCitizenship(citizenships []*model.Citizenship) []*model.Citizenship {
	var newCitizenships []*model.Citizenship
	for _, citizenship := range citizenships {
		newCitizenships = append(newCitizenships, citizenship)
	}
	return newCitizenships
}
