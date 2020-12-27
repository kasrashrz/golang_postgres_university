package models

import "gorm.io/gorm"

type University struct {
	gorm.Model
	Name         string   `json:"name"`
	Address      string   `json:"address"`
	URL 		 string	  `json:"url"`
    CreationDate string	  `json:"CreationDate"`
	Student []Student
	UniversityBranch []UniversityBranch
}