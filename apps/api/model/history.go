package model

import (
	"github.com/google/uuid"
	"time"
)

type History struct {
	Id             uuid.UUID `json:"Id"             gorm:"primary_key; column:Id; not null; type:char(36);"`
	CustomerId     uuid.UUID `json:"CustomerId"     gorm:"column:CustomerId; not null; type:char(36);"`
	Date           time.Time `json:"Date"           gorm:"column:Date; not null; index"`
	NumberOfPeople int       `json:"NumberOfPeople" gorm:"column:NumberOfPeople; not null"`
	Price          int       `json:"Price"          gorm:"column:Price; not null"`
	Note           string    `json:"Note"           gorm:"column:Note"`
	Room           string    `json:"Room"           gorm:"column:Room; not null"`
}
