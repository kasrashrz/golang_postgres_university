package models

type Students struct {
	Id           string `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Mail         string `json:"mail"`
	NationalCode string `json:"nationalCode"`
	Address      string `json:"address"`
	// Course       []Course `json:"course" `
}
type CreateStudentInput struct {
	Id           string `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Mail         string `json:"mail"`
	NationalCode string `json:"nationalCode"`
	Address      string `json:"address"`
}

type UpdateBookInput struct {
	Id           string `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Mail         string `json:"mail"`
	NationalCode string `json:"nationalCode"`
	Address      string `json:"address"`
}
