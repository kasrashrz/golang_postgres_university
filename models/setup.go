package models

import (
	"fmt"
	_ "fmt"
	"gorm.io/driver/postgres"
	//"log"

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
	db.AutoMigrate(&University{},&UniversityBranch{},&Student{},&Course{})
	//test_stu := Student{
	//	Name:         "xxxxx",
	//	Age:          10,
	//	Mail:         "none",
	//	NationalCode: "none",
	//	Address:      "teh",
	//	Courses:      []Course{
	//		{
	//			//Model:         gorm.Model{},
	//			Name:          "aljebra",
	//			QuantityPlace: 1,
	//			StartDate:     "1400",
	//			EndDate:       "1401",
	//			CreatedDate:   "1399",
	//		},
	//	},
	//
	//c1 := Course{
	//	Name:             "aljs ebra",
	//	QuantityPlace:    20,
	//	StartDate:        "1400",
	//	EndDate:          "1401",
	//	CreatedDate:      "1399",
	//	UniversityBranch: nil,
	//	Students:		  nil,
	//}

	//test_stu2 := Student{
	//	Name:         "mmdHasan",
	//	Age:          10,
	//	Mail:         "none",
	//	NationalCode: "none",
	//	Address:      "teh",
	//	Courses:      []Course{
	//		{
	//			//Model:         gorm.Model{},
	//			Name:          "math",
	//			QuantityPlace: 1,
	//			StartDate:     "1400",
	//			EndDate:       "1401",
	//			CreatedDate:   "1399",
	//		},
	//	},
	//}

    //uni := University{
	//	Name:             "azad",
	//	Address:          "teh",
	//	URL:              "azad.ir",
	//	CreationDate:     "1400",
	//	UniversityBranch: []UniversityBranch{{
	//		Name:         "azad_jonob",
	//		Address:      "jonob",
	//		URL:          "azdjnb.ir",
	//		CreationDate: "1400",
	//		UniversityID: 1,
	//		StudentID:    0,
	//	}},
	//}
	//
	//err = db.Save(&uni).Error
	//if err != nil {
	//	log.Fatal(err)
	//}
	//db.Save(uni)
	//db.(&Student{}, &Course{},&UniversityBranch{},&University{})
	//db.CreateTable(&Student{}, &Course{},&UniversityBranch{},&University{})

	fmt.Println("Done ")
	return db
}

