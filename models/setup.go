package models

import (
	"fmt"
	_ "fmt"
	"gorm.io/driver/postgres"
	"log"

	//"log"

	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
	//_ "github.com/lib/pq"
)

func SetupModels() *gorm.DB {
	dsn := "user=admin" + " password=root"  + " dbname=go_uni	 port=5432 TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,
	//})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&University{},&UniversityBranch{},&Student{},&Course{},&Teacher{})
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
	//	Name:             "aljebra",
	//	QuantityPlace:    10,
	//	StartDate:        "1400",
	//	EndDate:          "1401",
	//	CreatedDate:      "1399",
	//	UniversityBranches: nil,
	//	Students:		  nil,
	//}

	//stu := Student{
	//	Name:         "kasra",
	//	Age:          10,
	//	Mail:         "none",
	//	NationalCode: "0150553692",
	//	Address:      "teh",
	//	Courses:      []Course{
	//		{
	//			Model : gorm.Model{ID: 2},
	//		},
	//	},
	//	Teachers: nil,
	//}
	uniBranch := UniversityBranch{
		Name:         "azad jonoob",
		Address:      "tehran",
		URL:          "www.azad.ir",
		CreationDate: "none",
		Courses: []Course{
			{
				Model:gorm.Model{ID: 1},
			},
			{
				Model:gorm.Model{ID: 2},
			},
		},
	}
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
	err = db.Save(&uniBranch).Error
	if err != nil {
		log.Fatal(err)
	}
	//db.Save(uni)

	fmt.Println("Done ")
	return db
}

