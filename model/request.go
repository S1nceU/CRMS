package model

type CustomerRequest struct {
	CustomerId  int    `json:"CustomerId"`
	Name        string `json:"Name"`
	Gender      string `json:"Gender"`
	Birthday    string `json:"Birthday"`
	ID          string `json:"ID"`
	Address     string `json:"Address"`
	PhoneNumber string `json:"PhoneNumber"`
	CarNumber   string `json:"CarNumber"`
	Citizenship string `json:"Citizenship"`
	Note        string `json:"Note"`
}

type HistoryRequest struct {
	HistoryId      int    `json:"HistoryId"`
	CustomerId     int    `json:"CustomerId"`
	Date           string `json:"Date"`
	NumberOfPeople int    `json:"NumberOfPeople"`
	Price          int    `json:"Price"`
	Room           string `json:"Room"`
	Note           string `json:"Note"`
}
