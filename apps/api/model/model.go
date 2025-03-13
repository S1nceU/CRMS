package model

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	Id            uuid.UUID   `json:"Id"            gorm:"primary_key; column:Id; not null; type:char(36);"`
	Name          string      `json:"Name"          gorm:"column:Name; not null"`
	Gender        string      `json:"Gender"        gorm:"column:Gender; not null"`
	Birthday      time.Time   `json:"Birthday"      gorm:"column:Birthday; not null; index"`
	NationalId    string      `json:"NationalId"    gorm:"column:NationalId; not null; type:varchar(100); uniqueIndex;"`
	Address       string      `json:"Address"       gorm:"column:Address"`
	PhoneNumber   string      `json:"PhoneNumber"   gorm:"column:PhoneNumber"`
	CarNumber     string      `json:"CarNumber"     gorm:"column:CarNumber"`
	CitizenshipId int         `json:"CitizenshipId" gorm:"column:CitizenshipId; not null"`
	Citizenship   Citizenship `                     gorm:"foreignKey:CitizenshipId; references:Id"`
	Note          string      `json:"Note"          gorm:"column:Note"`
	Histories     []History   `                     gorm:"foreignKey:CustomerId; references:Id"`
}

type History struct {
	Id             uuid.UUID `json:"Id"             gorm:"primary_key; column:Id; not null; type:char(36);"`
	CustomerId     uuid.UUID `json:"CustomerId"     gorm:"column:CustomerId; not null; type:char(36);"`
	Date           time.Time `json:"Date"           gorm:"column:Date; not null; index"`
	NumberOfPeople int       `json:"NumberOfPeople" gorm:"column:NumberOfPeople; not null"`
	Price          int       `json:"Price"          gorm:"column:Price; not null"`
	Note           string    `json:"Note"           gorm:"column:Note"`
	Room           string    `json:"Room"           gorm:"column:Room; not null"`
}

type User struct {
	Id       uuid.UUID `json:"Id"       gorm:"primary_key; column:Id; not null; type:varchar(36);"`
	Username string    `json:"Username" gorm:"column:Username; not null; type:varchar(100); uniqueIndex;"`
	Password string    `json:"Password" gorm:"column:Password; not null; type:varchar(100);"`
}

type Citizenship struct {
	Id     int    `json:"Id"            gorm:"primary_key; column:Id; not null;"`
	Nation string `json:"Nation"        gorm:"column:Nation; not null; type:varchar(20); uniqueIndex;"`
	Alpha3 string `json:"Alpha3"        gorm:"column:Alpha3; not null; type:varchar(3); uniqueIndex;"`
}
