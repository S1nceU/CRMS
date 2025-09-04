package service

import (
	"errors"
	"time"

	"github.com/S1nceU/CRMS/apps/api/domain"
	"github.com/S1nceU/CRMS/apps/api/model"
	"github.com/google/uuid"
)

type HistoryService struct {
	repo domain.HistoryRepository
}

func NewHistoryService(repo domain.HistoryRepository) domain.HistoryService {
	return &HistoryService{
		repo: repo,
	}
}

func (u *HistoryService) ListHistories() ([]*model.History, error) {
	var err error
	var point []*model.History
	if point, err = u.repo.ListHistories(); err != nil {
		return nil, err
	}
	return convertToSliceOfHistory(point), err
}

func (u *HistoryService) ListHistoriesByCustomerId(in uuid.UUID) ([]*model.History, error) {
	var err error
	var point []*model.History
	newHistory := &model.History{
		CustomerId: in,
	}
	newCustomer := &model.Customer{
		Id: in,
	}
	if newCustomer, err = u.repo.ConfirmCustomerExistence(newCustomer); err != nil {
		return nil, err
	} else if newCustomer.Name == "" {
		return nil, errors.New("error CRMS : There is no this customer")
	}

	if point, err = u.repo.ListHistoriesByCustomer(newHistory); err != nil {
		return nil, err
	} else if len(point) == 0 {
		return nil, nil
	}
	return convertToSliceOfHistory(point), err
}

func (u *HistoryService) ListHistoriesForDate(in string) ([]*model.History, error) {
	var err error
	var point []*model.History
	var date time.Time
	if date, err = time.ParseInLocation("2006-01-02", in, time.Local); err != nil {
		return nil, errors.New("error CRMS : Date is incomplete")
	}
	if date.After(time.Now()) {
		return nil, errors.New("error CRMS : Date is after today")
	}
	newHistory := &model.History{
		Date: date,
	}
	if point, err = u.repo.ListHistoriesForDate(newHistory); err != nil {
		return nil, err
	} else if len(point) == 0 {
		return nil, errors.New("error CRMS : There was no customer in " + date.Format("2006-01-02"))
	}

	return convertToSliceOfHistory(point), err
}

func (u *HistoryService) ListHistoriesForDuring(in1 string, in2 string) ([]*model.History, error) {
	var err error
	var point []*model.History
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
	if point, err = u.repo.ListHistoriesForDuring(newHistory1, newHistory2); len(point) == 0 {
		return nil, errors.New("error CRMS : There is not any history between " + date1.Format("2006-01-02") + " to " + date2.Format("2006-01-02"))
	} else if err != nil {
		return nil, err
	}

	return convertToSliceOfHistory(point), err
}

func (u *HistoryService) GetHistoryByHistoryId(in uuid.UUID) (*model.History, error) {
	var err error
	newHistory := &model.History{
		Id: in,
	}
	if newHistory, err = u.repo.GetHistoryByHistoryId(newHistory); err != nil {
		return nil, err
	} else if newHistory.CustomerId == uuid.Nil {
		return nil, errors.New("error CRMS : There is no this history")
	}
	return newHistory, err
}

func (u *HistoryService) CreateHistory(in *model.History) (*model.History, error) {
	var err error
	var newHistory *model.History
	newCustomer := &model.Customer{
		Id: in.CustomerId,
	}
	if newCustomer, err = u.repo.ConfirmCustomerExistence(newCustomer); err != nil {
		return nil, err
	}

	if newCustomer.Name == "" {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	if err = validateHistoryInfo(in); err != nil {
		return nil, err
	}

	if _, err = u.repo.ListHistoriesByCustomer(in); err != nil {
		return nil, err
	}
	in.Id = uuid.New()
	if newHistory, err = u.repo.CreateHistory(in); err != nil {
		return nil, err
	}
	return newHistory, err
}

func (u *HistoryService) UpdateHistory(in *model.History) (*model.History, error) {
	var err error
	var newHistory *model.History
	newCustomer := &model.Customer{
		Id: in.CustomerId,
	}
	if newCustomer, err = u.repo.ConfirmCustomerExistence(newCustomer); err != nil {
		return nil, err
	}
	if newHistory, err = u.GetHistoryByHistoryId(in.Id); err != nil {
		return nil, err
	}

	if newCustomer.Name == "" {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	if newHistory.CustomerId == uuid.Nil {
		return nil, errors.New("error CRMS : There is no this history")
	}
	if err = validateHistoryInfo(in); err != nil {
		return nil, err
	}
	if _, err = u.GetHistoryByHistoryId(in.Id); err != nil {
		return nil, err
	}
	if newHistory, err = u.repo.UpdateHistory(in); err != nil {
		return nil, err
	}

	return newHistory, err
}

func (u *HistoryService) DeleteHistory(in uuid.UUID) error {
	var err error
	newHistory := &model.History{
		Id: in,
	}
	if _, err = u.GetHistoryByHistoryId(newHistory.Id); err != nil {
		return err
	}
	if err = u.repo.DeleteHistory(newHistory); err != nil {
		return err
	}
	return nil
}

func (u *HistoryService) DeleteHistoriesByCustomer(in uuid.UUID) error {
	var err error
	newHistory := &model.History{
		CustomerId: in,
	}
	if _, err = u.ListHistoriesByCustomerId(newHistory.CustomerId); err != nil {
		return err
	}
	if err = u.repo.DeleteHistoriesByCustomer(newHistory); err != nil {
		return err
	}
	return nil
}

func convertToSliceOfHistory(histories []*model.History) []*model.History {
	var historiesSlice []*model.History
	for _, history := range histories {
		historiesSlice = append(historiesSlice, history)
	}
	return historiesSlice
}

func validateHistoryInfo(history *model.History) error {
	if history.CustomerId == uuid.Nil {
		return errors.New("error CRMS : HistoryService Info is incomplete")
	}
	if history.Date.IsZero() {
		return errors.New("error CRMS : HistoryService Info is incomplete")
	}
	if history.NumberOfPeople == 0 {
		return errors.New("error CRMS : HistoryService Info is incomplete")
	}
	if history.Price == 0 {
		return errors.New("error CRMS : HistoryService Info is incomplete")
	}
	return nil
}
