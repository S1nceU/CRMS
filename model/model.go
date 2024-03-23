package model

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	CustomerId    uuid.UUID `json:"CustomerId"    gorm:"primary_key; column:CustomerId; not null; type:varchar(36);"`
	Name          string    `json:"Name"          gorm:"column:Name; not null"`
	Gender        string    `json:"Gender"        gorm:"column:Gender; not null"`
	Birthday      time.Time `json:"Birthday"      gorm:"column:Birthday; not null"`
	ID            string    `json:"ID"            gorm:"column:ID; not null; type:varchar(100); uniqueIndex; "`
	Address       string    `json:"Address"       gorm:"column:Address"`
	PhoneNumber   string    `json:"PhoneNumber"   gorm:"column:PhoneNumber"`
	CarNumber     string    `json:"CarNumber"     gorm:"column:CarNumber"`
	Note          string    `json:"Note"          gorm:"column:Note"`
	CitizenshipId int       `json:"Citizenship"   gorm:"column:CitizenshipId; not null"`
	History       []History `                     gorm:"foreignKey:CustomerId"`
}

type History struct {
	HistoryId      uuid.UUID `json:"HistoryId"      gorm:"primary_key; column:HistoryId; not null; type:varchar(36);"`
	CustomerId     uuid.UUID `json:"CustomerId"     gorm:"column:CustomerId; not null; type:varchar(36);"`
	Date           time.Time `json:"Date"           gorm:"column:Date; not null"`
	NumberOfPeople int       `json:"NumberOfPeople" gorm:"column:NumberOfPeople; not null"`
	Price          int       `json:"Price"          gorm:"column:Price; not null"`
	Note           string    `json:"Note"           gorm:"column:Note"`
	Room           string    `json:"Room"           gorm:"column:Room; not null"`
}

type User struct {
	UserId   uuid.UUID `json:"UserId"   gorm:"primary_key; column:UserId; not null; type:varchar(36);"`
	Username string    `json:"Username" gorm:"column:Username; not null; type:varchar(100); uniqueIndex;"`
	Password string    `json:"Password" gorm:"column:Password; not null; type:varchar(100);"`
}

type Citizenship struct {
	CitizenshipId int    `json:"CitizenshipId" gorm:"primary_key; column:CitizenshipId; not null;"`
	Nation        string `json:"Nation"        gorm:"column:Nation; not null; type:varchar(20); uniqueIndex;"`
	Alpha3        string `json:"Alpha3"        gorm:"column:Alpha3; not null; type:varchar(3); uniqueIndex;"`
}
