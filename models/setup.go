package models

import (
	"fmt"

	"gorm.io/gorm"

	// "github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
)

func SetupModels() *gorm.DB {
	dsn := "postgres://admin:root@127.0.0.1:5432/go_uni"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Students{})
	// m := Students{Id: "secondUser", Name: "arsalan", Age: 24, Mail: "a@mail", NationalCode: "fodaf", Address: "teh"}
	// db.Create(&m)
	return db
}
