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
