package models

import "github.com/jinzhu/gorm"
type Course struct {
	gorm.Model
	Name          string     `json:"name"`
	QuantityPlace int        `json:"name"`
	StartDate     string     `json:"startDate"`
	EndDate       string     `json:"EndDate"`
	CreatedDate   string     `json:"CreatedDate"`
	Student       []*Student `json:"students" gorm:"many2many:student_courses"`
}
