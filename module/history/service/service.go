package service

import (
	"errors"
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/history"
	"time"
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

func (u *HistoryService) GetHistoriesForDate(in time.Time) ([]model.History, error) {
	var err error
	var point []*model.History
	var out []model.History
	newHistory := &model.History{
		Date: in,
	}
	if point, err = u.repo.GetHistoriesForDate(newHistory); err != nil {
		return nil, err
	} else if len(point) == 0 {
		return nil, errors.New("error CRMS : There was no customer in " + in.Format("2006-01-02"))
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
	if in.Date.IsZero() {
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
	if in.Date.IsZero() {
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

func (u *HistoryService) GetHistoryForDuring(in1 string, in2 string) ([]model.History, error) {
	var err error
	var point []*model.History
	var out []model.History
	var date1 time.Time
	var date2 time.Time

	if date1, err = time.ParseInLocation("2006-01-02", in1, time.Local); err != nil {
		return nil, errors.New("error CRMS : Date is incomplete")
	}
	if date2, err = time.ParseInLocation("2006-01-02", in2, time.Local); err != nil {
		return nil, errors.New("error CRMS : Date is incomplete")
	}

	if date1.After(date2) {
		return nil, errors.New("error CRMS : Start date is after end date")
	}
	if date2.After(time.Now()) {
		return nil, errors.New("error CRMS : End date is after today")
	}

	date2 = date2.Add(time.Hour * 24)

	newHistory1 := &model.History{
		Date: date1,
	}
	newHistory2 := &model.History{
		Date: date2,
	}
	if point, err = u.repo.GetHistoryForDuring(newHistory1, newHistory2); len(point) == 0 {
		return nil, errors.New("error CRMS : There is not any history between " + date1.Format("2006-01-02") + " to " + date2.Format("2006-01-02"))
	} else if err != nil {
		return nil, err
	}

	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}
