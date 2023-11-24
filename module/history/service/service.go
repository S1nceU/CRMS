package service

import (
	"errors"
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/history"
)

type HistoryService struct {
	repo history.Repository
}

func NewHistoryService(repo history.Repository) history.Service {
	return &HistoryService{
		repo: repo,
	}
}

func (u *HistoryService) ListHistories() ([]model.History, error) {
	var err error
	var point []*model.History
	var out []model.History
	if point, err = u.repo.GetAllHistories(); err != nil {
		return nil, err
	}
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *HistoryService) GetHistoryByID(in int) ([]model.History, error) {
	var err error
	var point []*model.History
	var out []model.History
	newHistory := &model.History{
		CustomerId: in,
	}
	newCustomer := &model.Customer{
		CustomerId: in,
	}
	if newCustomer, err = u.repo.ConfirmCustomerExistence(newCustomer); err != nil {
		return nil, err
	} else if newCustomer.Name == "" {
		return nil, errors.New("error CRMS : There is no this customer")
	}

	if point, err = u.repo.GetHistoriesByCustomer(newHistory); err != nil {
		return nil, err
	} else if len(point) == 0 {
		return nil, errors.New("error CRMS : There is not any history")
	}
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *HistoryService) GetHistoriesForDate(in string) ([]model.History, error) {
	var err error
	var point []*model.History
	var out []model.History
	newHistory := &model.History{
		Date: in,
	}
	if point, err = u.repo.GetHistoriesForDate(newHistory); err != nil {
		return nil, err
	} else if len(point) == 0 {
		return nil, errors.New("error CRMS : There was no customer in " + in)
	}
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *HistoryService) GetHistoryByHistoryId(in int) (*model.History, error) {
	var err error
	newHistory := &model.History{
		HistoryId: in,
	}
	if newHistory, err = u.repo.GetHistoryByHistoryId(newHistory); err != nil {
		return nil, err
	} else if newHistory.CustomerId == 0 {
		return nil, errors.New("error CRMS : There is no this history")
	}
	return newHistory, err
}

func (u *HistoryService) CreateHistory(in *model.History) (*model.History, error) {
	var err error
	var newHistory *model.History
	newCustomer := &model.Customer{
		CustomerId: in.CustomerId,
	}
	if newCustomer, err = u.repo.ConfirmCustomerExistence(newCustomer); err != nil {
		return nil, err
	}

	if newCustomer.Name == "" {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	if in.CustomerId == 0 {
		return nil, errors.New("error CRMS : History Info is incomplete")
	}
	if in.Date == "" {
		return nil, errors.New("error CRMS : History Info is incomplete")
	}
	if in.NumberOfPeople == 0 {
		return nil, errors.New("error CRMS : History Info is incomplete")
	}
	if in.Price == 0 {
		return nil, errors.New("error CRMS : History Info is incomplete")
	}

	if _, err = u.repo.GetHistoriesByCustomer(in); err != nil {
		return nil, err
	}
	if newHistory, err = u.repo.CreateHistory(in); err != nil {
		return nil, err
	}
	return newHistory, err
}

func (u *HistoryService) UpdateHistory(in *model.History) (*model.History, error) {
	var err error
	var newHistory *model.History
	newCustomer := &model.Customer{
		CustomerId: in.CustomerId,
	}
	if newCustomer, err = u.repo.ConfirmCustomerExistence(newCustomer); err != nil {
		return nil, err
	}
	if newHistory, err = u.GetHistoryByHistoryId(in.HistoryId); err != nil {
		return nil, err
	}

	if newCustomer.Name == "" {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	if newHistory.CustomerId == 0 {
		return nil, errors.New("error CRMS : There is no this history")
	}
	if in.CustomerId == 0 {
		return nil, errors.New("error CRMS : History Info is incomplete")
	}
	if in.Date == "" {
		return nil, errors.New("error CRMS : History Info is incomplete")
	}
	if in.NumberOfPeople == 0 {
		return nil, errors.New("error CRMS : History Info is incomplete")
	}
	if in.Price == 0 {
		return nil, errors.New("error CRMS : History Info is incomplete")
	}
	if _, err = u.GetHistoryByHistoryId(in.HistoryId); err != nil {
		return nil, err
	}
	if newHistory, err = u.repo.UpdateHistory(in); err != nil {
		return nil, err
	}

	return newHistory, err
}

func (u *HistoryService) DeleteHistory(in int) error {
	var err error
	newHistory := &model.History{
		HistoryId: in,
	}
	if _, err = u.GetHistoryByHistoryId(newHistory.HistoryId); err != nil {
		return err
	}
	if err = u.repo.DeleteHistory(newHistory); err != nil {
		return err
	}
	return nil
}

func (u *HistoryService) DeleteHistoriesByCustomer(in int) error {
	var err error
	newHistory := &model.History{
		CustomerId: in,
	}
	if _, err = u.GetHistoryByID(newHistory.CustomerId); err != nil {
		return err
	}
	if err = u.repo.DeleteHistoriesByCustomer(newHistory); err != nil {
		return err
	}
	return nil
}
