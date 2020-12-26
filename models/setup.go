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

	test_stu := &Student{
		Name:         "ali",
		Age:          10,
		Mail:         "ali.maik",
		NationalCode: "fdoisdjvoijv0eiwja",
		Address:      "teh",
		Courses:      []Course{
			{
				Model:         gorm.Model{},
				Name:          "math",
				QuantityPlace: 1,
				StartDate:     "1400",
				EndDate:       "1401",
				CreatedDate:   "1399",
				Students:      nil,
			},
		},
	}
	//c1 := Course{
	//	Name:          "aljebra",
	//	QuantityPlace: 20,
	//	StartDate:     "1400",
	//	EndDate:       "1401",
	//	CreatedDate:   "1399",
	//	Students:      []Student{
	//		{
	//			Model:        gorm.Model{},
	//			Name:         "kasra",
	//			Age:          10,
	//			Mail:         "k.gamil",
	//			NationalCode: "asdfsdfsdfsd",
	//			Address:      "velenjak",
	//		},
	//	},
	//}
	//db.Save(&c1)
	db.Save(&test_stu)
	//db.DropTableIfExists(&Student{}, &Course{})
	//db.CreateTable(&Student{}, &Course{})

	return db
}
