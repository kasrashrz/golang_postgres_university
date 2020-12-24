package models

import (
	"fmt"

	"gorm.io/gorm"

	// "github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
)

func SetupModels() *gorm.DB {
	dsn := "postgres://admin:root@127.0.0.1:5432/go_uni"
	// dsn := "host=localhost user=admin password=root dbname=go_university port=9920"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Students{})
	if err != nil {
		fmt.Println(err)
	}
	m := Students{Id: "asdas!", address: "teh", name: "x", age: 18, mail: "sd@gmail.com", national_code: "58"}
	db.Create(&m)
	return db

}
