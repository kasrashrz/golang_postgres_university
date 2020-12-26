package models

import (
	"fmt"
	_ "fmt"

	//"gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func SetupModels() *gorm.DB {
	//dsn := "postgres://admin:root@127.0.0.1:5432/go_uni"
	//db, err := gorm.Open("postgres://admin:root@127.0.0.1:5432/go_uni")
	db, err := gorm.Open("postgres", "user=admin password=root dbname=go_uni ")
	if err != nil {
		panic(err)
	}
	//err := db.DB().Ping()
	fmt.Print(err)

	test_stu := Student{
		Name:         "aliv",
		Age:          10,
		Mail:         "asdaasdsdasd",
		NationalCode: "aasasddasdas",
		Address:      "asdasadsdasd",
		CouseID:      uint(1),
		Courses:      []Course{},
	}
	c1 := Course{
		Name:          "math",
		QuantityPlace: 20,
		StartDate:     "1400",
		EndDate:       "1401",
		CreatedDate:   "1399",
		StudentID:     uint(1),
		Students:      []Student{},
	}
	db.Save(&test_stu)
	db.Save(&c1)
	// db.Create(&c1)
	// db.DropTableIfExists(&Student{}, &Course{})
	// db.CreateTable(&Student{}, &Course{})

	return db
}
