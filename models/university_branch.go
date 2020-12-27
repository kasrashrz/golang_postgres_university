package models

import "gorm.io/gorm"

type UniversityBranch struct {
	gorm.Model
	Name         string   `json:"name"`
	Address      string   `json:"address"`
	URL 		 string	  `json:"url"`
	CreationDate string	  `json:"CreationDate"`
	UniversityID uint	  `json:"university_id"`
	Courses      []Course `gorm:"many2many:UniversityBranch_courses"`
}
