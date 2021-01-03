package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Mail         string   `json:"mail"`
	NationalCode string   `gorm:"unique"json:"nationalCode"`
	Address      string   `json:"address"`
	Courses      []Course `gorm:"many2many:students_courses;"`
	//Teachers      []Teacher `gorm:"many2many:students_teachers;"`
	UniversityID	uint 	`json:"university_id"`
	University University `gorm:"references:ID"`
}
//
//type CreateStudentInput struct {
//	gorm.Model
//	Name         string   `json:"name"`
//	Age          int      `json:"age"`
//	Mail         string   `json:"mail"`
//	NationalCode string   `gorm:"unique"json:"nationalCode"`
//	Address      string   `json:"address"`
//	Courses      []Course `gorm:"many2many:students_courses;"`
//	//Teachers      []Teacher `gorm:"many2many:students_teachers;"`
//	UniversityID	uint 	`json:"university_id"`
//	University University `gorm:"references:ID"`
//}
//
//type UpdateStudentInput struct {
//	gorm.Model
//	Name         string   `json:"name"`
//	Age          int      `json:"age"`
//	Mail         string   `json:"mail"`
//	NationalCode string   `gorm:"unique"json:"nationalCode"`
//	Address      string   `json:"address"`
//	Courses      []Course `gorm:"many2many:students_courses;"`
//	//Teachers      []Teacher `gorm:"many2many:students_teachers;"`
//	UniversityID	uint 	`json:"university_id"`
//	University University `gorm:"references:ID"`
//}
