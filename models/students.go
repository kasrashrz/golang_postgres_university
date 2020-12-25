package models

import "github.com/jinzhu/gorm"
type Student struct {
	gorm.Model
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Mail         string `json:"mail"`
	NationalCode string `json:"nationalCode"`
	Address      string `json:"address"`
	Courses       []*Course `json:"course" gorm:"many2many:student_courses"`
}
type CreateStudentInput struct {
	gorm.Model
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Mail         string `json:"mail"`
	NationalCode string `json:"nationalCode"`
	Address      string `json:"address"`
}

type UpdateBookInput struct {
	gorm.Model
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Mail         string `json:"mail"`
	NationalCode string `json:"nationalCode"`
	Address      string `json:"address"`
}
