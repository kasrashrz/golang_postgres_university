package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name          string  `json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `gorm:"unique;" json:"nationalCode"`
	//CourseID uint
	//Course      []Course    `gorm:"references:ID"`
}
type CreateTeacherInput struct {
	gorm.Model
	Name          string    `json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `gorm:"unique;" json:"nationalCode"`
	//CourseID uint
	//Course      []Course  `gorm:"references:ID"`
}
type UpdateTeacherInput struct {
	gorm.Model
	Name          string    `json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `gorm:"unique;" json:"nationalCode"`
	//CourseID uint
	//Course      []Course  `gorm:"references:ID"`
}