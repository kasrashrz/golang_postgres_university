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
	//}

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
	//	Name:         "azad uni",
	//	Address:      "none",
	//	URL:          "none",
	//	CreationDate: "none",
	//	Student: []Student{{
	//		Name:         "ali",
	//		Age:          10,
	//		Mail:         "ali.maik",
	//		NationalCode: "fdoisdjvoijv0eiwja",
	//		Address:      "teh",
	//		Courses:      nil,
	//		UniversityID: 1,
	//	}},
	//	UniversityBranch: []UniversityBranch{{ Name: "azad shomal", Address:"darabad" , URL: "www.Ashomal.ir", CreationDate: "1399", UniversityID: 1 , Courses: []Course{{ Name: "chemistry", QuantityPlace: 10 , StartDate: "1400" , EndDate:"1400" , CreatedDate:"1400"  }} }},
	//}
	uni := UniversityBranch{
		Name:         "azad jonob",
		Address:      "gisha",
		URL:          "www.azadjonob.ir",
		CreationDate: "none",
		UniversityID: 3,
		Courses: []Course{{
			Name:             "math",
			QuantityPlace:    1,
			StartDate:        "1400",
			EndDate:          "1401",
			CreatedDate:      "1399",
			//UniversityBranch: nil,
			//Students:         nil,
		}},
	}
	//db.Save(&c1)
	//err = db.Save(&test_stu).Error
	//if err != nil {
	//	log.Fatal(err)
	//}
	err = db.Save(&uni).Error
	if err != nil {
		log.Fatal(err)
	}
	//db.Save(uni)
	//db.(&Student{}, &Course{},&UniversityBranch{},&University{})
	//db.CreateTable(&Student{}, &Course{},&UniversityBranch{},&University{})
	fmt.Println("Done ")
	return db
}

