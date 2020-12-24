package models

type Students struct {
	Id            string `json:"id" gorm:"primary_key"`
	name          string `json:"name"`
	age           int    `json:"age"`
	mail          string `json:"mail"`
	national_code string `json:"nationalCode"`
	address       string `json:"address"`
}
