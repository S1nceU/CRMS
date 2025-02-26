package repository

import (
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	orm *gorm.DB
}

func NewHistoryRepository(orm *gorm.DB) domain.HistoryRepository {
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
	err := u.orm.Where("CustomerId = ?", history.CustomerId).Find(&histories).Error
	return histories, err
}

func (u *HistoryRepository) ListHistoriesForDate(history *model.History) ([]*model.History, error) {
	var histories []*model.History
	err := u.orm.Where("Date = ?", history.Date).Find(&histories).Error
	return histories, err
}

func (u *HistoryRepository) ListHistoriesForDuring(history1 *model.History, history2 *model.History) ([]*model.History, error) {
	var histories []*model.History
	err := u.orm.Where("Date >= ? AND Date <= ?", history1.Date, history2.Date).Find(&histories).Error
	return histories, err
}

func (u *HistoryRepository) GetHistoryByHistoryId(history *model.History) (*model.History, error) {
	err := u.orm.Where("Id = ?", history.Id).Find(&history).Error
	return history, err
}

func (u *HistoryRepository) CreateHistory(history *model.History) (*model.History, error) {
	err := u.orm.Create(&history).Error
	return history, err
}

func (u *HistoryRepository) UpdateHistory(history *model.History) (*model.History, error) {
	err := u.orm.Model(history).Where("Id = ?", history.Id).Updates(&history).Error
	return history, err
}

func (u *HistoryRepository) DeleteHistory(history *model.History) error {
	err := u.orm.Where("Id = ?", history.Id).Delete(&history).Error
	return err
}

func (u *HistoryRepository) DeleteHistoriesByCustomer(history *model.History) error {
	err := u.orm.Where("CustomerId = ?", history.CustomerId).Delete(&history).Error
	return err
}

func (u *HistoryRepository) ConfirmCustomerExistence(customer *model.Customer) (*model.Customer, error) {
	err := u.orm.Where("Id = ?", customer.Id).Find(&customer).Error
	return customer, err
}
