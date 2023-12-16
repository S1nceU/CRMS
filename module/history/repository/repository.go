package repository

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/history"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	orm *gorm.DB
}

func NewHistoryRepository(orm *gorm.DB) history.Repository {
	return &HistoryRepository{
		orm: orm,
	}
}

func (u *HistoryRepository) GetAllHistories() ([]*model.History, error) {
	var err error
	var in []*model.History
	err = u.orm.Find(&in).Error
	return in, err
}

func (u *HistoryRepository) GetHistoriesByCustomer(in *model.History) ([]*model.History, error) {
	var err error
	var out []*model.History
	err = u.orm.Where("customer_id = ?", in.CustomerId).Find(&out).Error
	return out, err
}

func (u *HistoryRepository) GetHistoriesForDate(in *model.History) ([]*model.History, error) {
	var err error
	var out []*model.History
	err = u.orm.Where("date = ?", in.Date).Find(&out).Error
	return out, err
}

func (u *HistoryRepository) GetHistoryByHistoryId(in *model.History) (*model.History, error) {
	var err error
	err = u.orm.Where("history_id = ?", in.HistoryId).Find(&in).Error
	return in, err
}

func (u *HistoryRepository) CreateHistory(in *model.History) (*model.History, error) {
	var err error
	err = u.orm.Create(&in).Error
	return in, err
}

func (u *HistoryRepository) UpdateHistory(in *model.History) (*model.History, error) {
	var err error
	err = u.orm.Model(in).Where("history_id = ?", in.HistoryId).Updates(&in).Error
	return in, err
}

func (u *HistoryRepository) DeleteHistory(in *model.History) error {
	var err error
	err = u.orm.Where("history_id = ?", in.HistoryId).Delete(&in).Error
	return err
}

func (u *HistoryRepository) ConfirmCustomerExistence(in *model.Customer) (*model.Customer, error) {
	var err error
	err = u.orm.Where("customer_id = ?", in.CustomerId).Find(&in).Error
	return in, err
}

func (u *HistoryRepository) DeleteHistoriesByCustomer(in *model.History) error {
	var err error
	err = u.orm.Where("customer_id = ?", in.CustomerId).Delete(&in).Error
	return err
}

func (u *HistoryRepository) GetHistoryForDuring(in1 *model.History, in2 *model.History) ([]*model.History, error) {
	var err error
	var out []*model.History
	err = u.orm.Where("date >= ? AND date <= ?", in1.Date, in2.Date).Find(&out).Error
	return out, err
}
