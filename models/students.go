package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Mail         string   `json:"mail"`
	NationalCode string   `json:"nationalCode"`
	Address      string   `json:"address"`
	Courses      []Course `gorm:"many2many:students_courses"`
	UniversityID	uint 	`json:"university_id"`
}

type CreateStudentInput struct {
	gorm.Model
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Mail         string   `json:"mail"`
	NationalCode string   `json:"nationalCode"`
	Address      string   `json:"address"`
	//CourseID     uint     `sql:"index"`
	Courses      []Course `gorm:"many2many:students_courses"`
}

type UpdateBookInput struct {
	gorm.Model
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Mail         string   `json:"mail"`
	NationalCode string   `json:"nationalCode"`
	Address      string   `json:"address"`
	//CourseID     uint     `sql:"index"`
	Courses      []Course `gorm:"many2many:students_courses"`
}
