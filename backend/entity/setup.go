package entity

import (
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("SE-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(&Course{},
		&Qualification{},
	)

	db = database

	//Qualification
	qualification1 := Qualification{
		Qualification_ID:   "QT01",
		Qualification_Name: "Bachelor's degree",
	}
	db.Model(&Qualification{}).Create(&qualification1)

	qualification2 := Qualification{
		Qualification_ID:   "QT02",
		Qualification_Name: "master's degree",
	}
	db.Model(&Qualification{}).Create(&qualification2)

	qualification3 := Qualification{
		Qualification_ID:   "QT03",
		Qualification_Name: "Ph.D.",
	}
	db.Model(&Qualification{}).Create(&qualification3)

	major1 := Major{
		Major_ID:   "CPE",
		Major_Name: "Computer Engineering",
	}
	db.Model(&Major{}).Create(&major1)

	major2 := Major{
		Major_ID:   "EE",
		Major_Name: "Electrical Engineering",
	}
	db.Model(&Major{}).Create(&major2)

	//Course
	course1 := Course{
		Course_ID:     "CPE2560",
		Course_Name:   "หลักสูตรวิศวกรรมคอมพิวเตอร์ 2560",
		Datetime:      time.Now(),
		Qualification: qualification1,
		Major:         major1,
	}
	db.Model(&Course{}).Create(&course1)

	course2 := Course{
		Course_ID:     "CPE2564",
		Course_Name:   "หลักสูตรวิศวกรรมคอมพิวเตอร์ 2564",
		Datetime:      time.Now(),
		Qualification: qualification2,
		Major:         major1,
	}
	db.Model(&Course{}).Create(&course2)

	course3 := Course{
		Course_ID:     "EE2560",
		Course_Name:   "หลักสูตรวิศวกรรมไฟฟ้า2560",
		Datetime:      time.Now(),
		Qualification: qualification3,
		Major:         major2,
	}
	db.Model(&Course{}).Create(&course3)

}
