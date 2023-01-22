package entity

import (
	"time"
)

type Qualification struct {
	Qualification_ID string `gorm:"primaryKey"`

	Qualification_Name string

	Courses []Course `gorm:"foreignKey:Qualification_ID"`
}

type Major struct {
	Major_ID string `gorm:"primaryKey"`

	Major_Name string

	Courses []Course `gorm:"foreignKey:Major_ID"`
}

type Course struct {
	Course_ID string `gorm:"primaryKey"`

	Course_Name string

	Datetime time.Time

	Qualification_ID *string
	Qualification    Qualification `gorm:"references:Qualification_ID"`

	Major_ID *string
	Major    Major `gorm:"references:Major_ID"`
}
