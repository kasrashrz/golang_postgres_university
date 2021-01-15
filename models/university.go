package models

import "gorm.io/gorm"

type University struct {
	gorm.Model
	Name         string   `json:"name"`
	Address      string   `json:"address"`
	URL 		 string	  `json:"url"`
    CreationDate string	  `json:"CreationDate"`
	//Students []Student 		`gorm:"ForeignKey:Name;"`
	//UniversityBranches []UniversityBranch
}
//
//type CreateUniversityInput struct {
//	gorm.Model
//	Name         string   `json:"name"`
//	Address      string   `json:"address"`
//	URL 		 string	  `json:"url"`
//	CreationDate string	  `json:"CreationDate"`
//	Students []Student
//	UniversityBranches []UniversityBranch
//}
//
//type UpdateUniversityInput struct {
//	gorm.Model
//	Name         string   `json:"name"`
//	Address      string   `json:"address"`
//	URL 		 string	  `json:"url"`
//	CreationDate string	  `json:"CreationDate"`
//	Students []Student
//	UniversityBranches []UniversityBranch
//}