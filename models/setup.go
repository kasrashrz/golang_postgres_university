package models

import (
	"fmt"

	"gorm.io/gorm"

	// "github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
)

func SetupModels() *gorm.DB {
	dsn := "host=localhost user=admin password=root dbname=go_university port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Students{})
	if err == nil {
		fmt.Println(err)
	}
	m := Students{Id: "ksa52ADj8ds2!", address: "teh"}
	db.Create(&m)
	return db
}
