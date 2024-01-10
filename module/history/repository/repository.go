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

func (u *HistoryRepository) ListHistories() ([]*model.History, error) {
	var histories []*model.History
	err := u.orm.Find(&histories).Error
	return histories, err
}

func (u *HistoryRepository) ListHistoriesByCustomer(history *model.History) ([]*model.History, error) {
	var histories []*model.History
	err := u.orm.Where("customer_id = ?", history.CustomerId).Find(&histories).Error
	return histories, err
}

func (u *HistoryRepository) ListHistoriesForDate(history *model.History) ([]*model.History, error) {
	var histories []*model.History
	err := u.orm.Where("date = ?", history.Date).Find(&histories).Error
	return histories, err
}

func (u *HistoryRepository) ListHistoriesForDuring(history1 *model.History, history2 *model.History) ([]*model.History, error) {
	var histories []*model.History
	err := u.orm.Where("date >= ? AND date <= ?", history1.Date, history2.Date).Find(&histories).Error
	return histories, err
}

func (u *HistoryRepository) GetHistoryByHistoryId(history *model.History) (*model.History, error) {
	err := u.orm.Where("history_id = ?", history.HistoryId).Find(&history).Error
	return history, err
}

func (u *HistoryRepository) CreateHistory(history *model.History) (*model.History, error) {
	err := u.orm.Create(&history).Error
	return history, err
}

func (u *HistoryRepository) UpdateHistory(history *model.History) (*model.History, error) {
	err := u.orm.Model(history).Where("history_id = ?", history.HistoryId).Updates(&history).Error
	return history, err
}

func (u *HistoryRepository) DeleteHistory(history *model.History) error {
	err := u.orm.Where("history_id = ?", history.HistoryId).Delete(&history).Error
	return err
}

func (u *HistoryRepository) DeleteHistoriesByCustomer(history *model.History) error {
	err := u.orm.Where("customer_id = ?", history.CustomerId).Delete(&history).Error
	return err
}

func (u *HistoryRepository) ConfirmCustomerExistence(customer *model.Customer) (*model.Customer, error) {
	err := u.orm.Where("customer_id = ?", customer.CustomerId).Find(&customer).Error
	return customer, err
}
