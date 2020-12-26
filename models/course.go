package models

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name          string    `gorm:"unique;" json:"name"`
	QuantityPlace int       `json:"quantityPlace"`
	StartDate     string    `json:"startDate"`
	EndDate       string    `json:"EndDate"`
	CreatedDate   string    `json:"CreatedDate"`
	//StudentID     uint      `sql:"index"`
	Students      []Student `gorm:"many2many:students_courses"`
}
