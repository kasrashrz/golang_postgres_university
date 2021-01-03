package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name          string    `json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `gorm:"unique;" json:"nationalCode"`
	Students      []Student `gorm:"many2many:students_teachers"`
}
type CreateTeacherInput struct {
	gorm.Model
	Name          string    `json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `gorm:"unique;" json:"nationalCode"`
	Courses      []Course `gorm:"many2many:students_courses"`
	Students      []Student `gorm:"many2many:students_teachers"`
}
type UpdateTeacherInput struct {
	gorm.Model
	Name          string    `json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `gorm:"unique;" json:"nationalCode"`
	Courses      []Course `gorm:"many2many:students_courses"`
	Students      []Student `gorm:"many2many:students_teachers"`
}