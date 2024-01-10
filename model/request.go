package model

import "github.com/google/uuid"

type CustomerRequest struct {
	CustomerId  uuid.UUID `json:"CustomerId"`
	Name        string    `json:"Name"`
	Gender      string    `json:"Gender"`
	Birthday    string    `json:"Birthday"`
	ID          string    `json:"ID"`
	Address     string    `json:"Address"`
	PhoneNumber string    `json:"PhoneNumber"`
	CarNumber   string    `json:"CarNumber"`
	Citizenship string    `json:"Citizenship"`
	Note        string    `json:"Note"`
}

type CustomerIdRequest struct {
	CustomerId uuid.UUID `json:"CustomerId"`
}

type HistoryRequest struct {
	HistoryId      uuid.UUID `json:"HistoryId"`
	CustomerId     uuid.UUID `json:"CustomerId"`
	Date           string    `json:"Date"`
	NumberOfPeople int       `json:"NumberOfPeople"`
	Price          int       `json:"Price"`
	Room           string    `json:"Room"`
	Note           string    `json:"Note"`
}

type DuringRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type DateRequest struct {
	Date string `json:"Date"`
}

type HistoryIdRequest struct {
	HistoryId uuid.UUID `json:"HistoryId"`
}
