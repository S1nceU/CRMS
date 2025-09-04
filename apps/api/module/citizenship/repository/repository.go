package repository

import (
	"github.com/S1nceU/CRMS/apps/api/domain"
	"github.com/S1nceU/CRMS/apps/api/model"
	"gorm.io/gorm"
)

type CitizenshipRepository struct {
	orm *gorm.DB
}

func NewCitizenshipRepository(orm *gorm.DB) domain.CitizenshipRepository {
	return &CitizenshipRepository{
		orm: orm,
	}
}

func (u *CitizenshipRepository) ListCitizenships() ([]*model.Citizenship, error) {
	var citizenships []*model.Citizenship
	err := u.orm.Find(&citizenships).Error
	return citizenships, err
}

func (u *CitizenshipRepository) GetCitizenshipByID(citizenship *model.Citizenship) (*model.Citizenship, error) {
	err := u.orm.Where("Id = ?", citizenship.Id).Find(&citizenship).Error
	return citizenship, err
}

func (u *CitizenshipRepository) GetCitizenshipByCitizenshipName(citizenship *model.Citizenship) (*model.Citizenship, error) {
	err := u.orm.Where("Nation = ?", citizenship.Nation).Find(&citizenship).Error
	return citizenship, err
}
