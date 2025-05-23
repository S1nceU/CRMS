package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"Id"       gorm:"primary_key; column:Id; not null; type:varchar(36);"`
	Username string    `json:"Username" gorm:"column:Username; not null; type:varchar(100); uniqueIndex;"`
	Password string    `json:"Password" gorm:"column:Password; not null; type:varchar(100);"`
}
