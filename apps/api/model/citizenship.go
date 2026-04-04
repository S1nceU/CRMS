package model

type Citizenship struct {
	Sequence int    `json:"Sequence"     gorm:"primaryKey; column:Sequence; not null; autoIncrement;"`
	Id       int    `json:"Id"            gorm:"column:Id; not null; uniqueIndex;"`
	Nation   string `json:"Nation"        gorm:"column:Nation; not null; type:varchar(20); uniqueIndex;"`
	Alpha3   string `json:"Alpha3"        gorm:"column:Alpha3; not null; type:varchar(3); uniqueIndex;"`
}

