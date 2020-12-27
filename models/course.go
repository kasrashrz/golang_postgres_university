package models

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name          string    `gorm:"unique;" json:"name"`
	QuantityPlace int       `json:"quantityPlace"`
	StartDate     string    `json:"startDate"`
	EndDate       string    `json:"EndDate"`
	CreatedDate   string    `json:"CreatedDate"`
	UniversityBranch	[]UniversityBranch `gorm:"many2many:UniversityBranch_courses"`
	Students      []Student `gorm:"many2many:students_courses"`
}

type CreateCourseInput struct {
	gorm.Model
	Name          string    `gorm:"unique;" json:"name"`
	QuantityPlace int       `json:"quantityPlace"`
	StartDate     string    `json:"startDate"`
	EndDate       string    `json:"EndDate"`
	CreatedDate   string    `json:"CreatedDate"`
	UniversityBranch	[]UniversityBranch `gorm:"many2many:UniversityBranch_courses"`
	Students      []Student `gorm:"many2many:students_courses"`
}

type UpdateCourseInput struct {
	gorm.Model
	Name          string    `gorm:"unique;" json:"name"`
	QuantityPlace int       `json:"quantityPlace"`
	StartDate     string    `json:"startDate"`
	EndDate       string    `json:"EndDate"`
	CreatedDate   string    `json:"CreatedDate"`
	UniversityBranch	[]UniversityBranch `gorm:"many2many:UniversityBranch_courses"`
	Students      []Student `gorm:"many2many:students_courses"`
}
