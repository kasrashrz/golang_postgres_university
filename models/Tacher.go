package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name          string    `gorm:"unique;"json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `json:"nationalCode"`
	Courses      []Course `gorm:"many2many:students_courses"`
	//UniversityBranch	[]UniversityBranch `gorm:"many2many:UniversityBranch_courses"`
	Students      []Student `gorm:"many2many:students_teachers"`
}

type CreateTeacherInput struct {
	gorm.Model
	Name          string    `gorm:"unique;"json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `json:"nationalCode"`
	Courses      []Course `gorm:"many2many:students_courses"`
	//UniversityBranch	[]UniversityBranch `gorm:"many2many:UniversityBranch_courses"`
	Students      []Student `gorm:"many2many:students_teachers"`
}
type UpdateTeacherInput struct {
	gorm.Model
	Name          string    `gorm:"unique;"json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `json:"nationalCode"`
	Courses      []Course `gorm:"many2many:students_courses"`
	//UniversityBranch	[]UniversityBranch `gorm:"many2many:UniversityBranch_courses"`
	Students      []Student `gorm:"many2many:students_teachers"`
}