package models

import "gorm.io/gorm"

type UniversityBranch struct {
	gorm.Model
	Name         string   `gorm:"unique;"json:"name"`
	Address      string   `json:"address"`
	URL 		 string	  `json:"url"`
	CreationDate string	  `json:"CreationDate"`
	UniversityID uint
	University University `gorm:"references:ID"`
	Courses      []Course `gorm:"many2many:UniversityBranch_courses"`
}
//
//type CreateUniversityBranchInput struct {
//	gorm.Model
//	Name         string   `gorm:"unique;"json:"name"`
//	Address      string   `json:"address"`
//	URL 		 string	  `json:"url"`
//	CreationDate string	  `json:"CreationDate"`
//	UniversityID uint
//	University University  `gorm:"references:ID"`
//	Courses      []Course `gorm:"many2many:UniversityBranch_courses"`
//}
//
//type UpdateUniversityBranchInput struct {
//	gorm.Model
//	Name         string   `gorm:"unique;"json:"name"`
//	Address      string   `json:"address"`
//	URL 		 string	  `json:"url"`
//	CreationDate string	  `json:"CreationDate"`
//	University University `gorm:"references:ID"`
//	Courses      []Course `gorm:"many2many:UniversityBranch_courses"`
//}
