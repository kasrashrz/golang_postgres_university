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
	db, err := gorm.Open("postgres","user=admin password=root dbname=go_uni ")
	if err != nil {
		panic(err)
	}
	//err := db.DB().Ping()
	fmt.Print(err)
	test_stu := Student{
		Name:         "sad",
		Age:          10,
		Mail:         "asdaasdsdasd",
		NationalCode: "aasasddasdas",
		Address:      "asdasadsdasd",
		Courses:      nil,
	}
	//c1 := Course{
	//	//Model:         gorm.Model{},
	//	Name:          "physic",
	//	QuantityPlace: 1,
	//	StartDate:     "1400",
	//	EndDate:       "1401",
	//	CreatedDate:   "1399",
	//	Student: []*Student{},
	//}
	//db.Create(&test_stu)
	//db.Create(&c1)
	db.Save(&test_stu)
	return db
}
