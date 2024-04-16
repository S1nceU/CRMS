package model

import "github.com/google/uuid"

// Customer Request

type CustomerNationalIdRequest struct {
	NationalId string `json:"NationalId"`
}

type CustomerRequest struct {
	CustomerId  uuid.UUID `json:"CustomerId"`
	Name        string    `json:"Name"`
	Gender      string    `json:"Gender"`
	Birthday    string    `json:"Birthday"`
	NationalId  string    `json:"NationalId"`
	Address     string    `json:"Address"`
	PhoneNumber string    `json:"PhoneNumber"`
	CarNumber   string    `json:"CarNumber"`
	Citizenship int       `json:"Citizenship"`
	Note        string    `json:"Note"`
}

type CustomerNameRequest struct {
	Name string `json:"Name"`
}

type CustomerIdRequest struct {
	CustomerId uuid.UUID `json:"CustomerId"`
}

type CustomerCitizenshipRequest struct {
	Citizenship int `json:"Citizenship"`
}

type CustomerPhoneRequest struct {
	PhoneNumber string `json:"PhoneNumber"`
}

// History Request

type HistoryRequest struct {
	HistoryId      uuid.UUID `json:"HistoryId"`
	CustomerId     uuid.UUID `json:"CustomerId"`
	Date           string    `json:"Date"`
	NumberOfPeople int       `json:"NumberOfPeople"`
	Price          int       `json:"Price"`
	Room           string    `json:"Room"`
	Note           string    `json:"Note"`
}

type HistoryIdRequest struct {
	HistoryId uuid.UUID `json:"HistoryId"`
}

type DuringRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type DateRequest struct {
	Date string `json:"Date"`
}

type HistoryCustomerIdRequest struct {
	CustomerId uuid.UUID `json:"CustomerId"`
}

// Citizenship Request

type CitizenshipRequest struct {
	CitizenshipId int `json:"CitizenshipId"`
}

type CitizenshipNameRequest struct {
	CitizenshipName string `json:"CitizenshipName"`
}

// User Request

type UserLoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type UserTokenRequest struct {
	Token string `json:"Token"`
}
