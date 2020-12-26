package models

import (
	"fmt"
	_ "fmt"
	"gorm.io/driver/postgres"
	"log"

	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
	//_ "github.com/lib/pq"
)

func SetupModels() *gorm.DB {
	dsn := "user=admin" + " password=root"  + " dbname=go_uni	 port=5432 TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Student{})
	test_stu := Student{
		Name:         "xxxxx",
		Age:          10,
		Mail:         "none",
		NationalCode: "none",
		Address:      "teh",
		Courses:      []Course{
			{
				//Model:         gorm.Model{},
				Name:          "aljebra",
				QuantityPlace: 1,
				StartDate:     "1400",
				EndDate:       "1401",
				CreatedDate:   "1399",
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
	//			//Model:        gorm.Model{},
	//			Name:         "khamenii",
	//			Age:          10,
	//			Mail:         "k.gamil",
	//			NationalCode: "asdfsdfsdfsd",
	//			Address:      "velenjak",
	//		},
	//	},
	//}

	test_stu2 := Student{
		Name:         "mmdHasan",
		Age:          10,
		Mail:         "none",
		NationalCode: "none",
		Address:      "teh",
		Courses:      []Course{
			{
				//Model:         gorm.Model{},
				Name:          "math",
				QuantityPlace: 1,
				StartDate:     "1400",
				EndDate:       "1401",
				CreatedDate:   "1399",
			},
		},
	}

	//db.Save(&c1)
	err = db.Save(&test_stu).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Save( &test_stu2).Error
	if err != nil {
		log.Fatal(err)
	}
	//db.DropTableIfExists(&Student{}, &Course{})
	//db.CreateTable(&Student{}, &Course{})
	fmt.Println("Done ")
	return db
}

