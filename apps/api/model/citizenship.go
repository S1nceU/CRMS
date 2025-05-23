package model

type Citizenship struct {
	Id     int    `json:"Id"            gorm:"primary_key; column:Id; not null;"`
	Nation string `json:"Nation"        gorm:"column:Nation; not null; type:varchar(20); uniqueIndex;"`
	Alpha3 string `json:"Alpha3"        gorm:"column:Alpha3; not null; type:varchar(3); uniqueIndex;"`
}

