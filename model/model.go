package model

import "time"

type Customer struct {
	CustomerId  int       `json:"CustomerId"  gorm:"primary_key; auto_increase;not null"`
	Name        string    `json:"Name"        gorm:"column:Name; not null"`
	Gender      string    `json:"Gender"      gorm:"column:Gender; not null"`
	Birthday    time.Time `json:"Birthday"    gorm:"column:Birthday; not null"`
	ID          string    `json:"ID"          gorm:"column:ID; type:varchar(100); uniqueIndex; not null"`
	Address     string    `json:"Address"     gorm:"column:Address"`
	PhoneNumber string    `json:"PhoneNumber" gorm:"column:PhoneNumber"`
	CarNumber   string    `json:"CarNumber"   gorm:"column:CarNumber"`
	Citizenship string    `json:"Citizenship" gorm:"column:Citizenship; not null"`
	Note        string    `json:"Note"        gorm:"column:Note"`
	History     History   `                   gorm:"foreignKey:CustomerId"`
}

type History struct {
	HistoryId      int       `json:"HistoryId"      gorm:"primary_key;auto_increase; not null"`
	CustomerId     int       `json:"CustomerId"     gorm:"column:Customer_id; not null"`
	Date           time.Time `json:"Date"           gorm:"column:Date; not null"`
	NumberOfPeople int       `json:"NumberOfPeople" gorm:"column:NumberOfPeople; not null"`
	Price          int       `json:"Price"          gorm:"column:Price; not null"`
	Room           string    `json:"Room"           gorm:"column:Room; not null"`
	Note           string    `json:"Note"           gorm:"column:Note"`
}
