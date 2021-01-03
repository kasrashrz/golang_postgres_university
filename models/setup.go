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

	stu := Student{
		Name:         "kasra",
		Age:         	18,
		Mail:         "kk@mailmail",
		NationalCode: "015asd",
		Address:      "teh",
		Courses: []Course{
			{
				Model:gorm.Model{ID: 1},
			},
		},
		UniversityID: 1,
		University:   University{},
	}
	//
	//uni := University{
	//	Name:         "azad",
	//	Address:      "tehran",
	//	URL:          "www.azad.com",
	//	CreationDate: "1400",
	//}
	//uni_branch := UniversityBranch{
	//	Name:         "azad jonob",
	//	Address:      "none",
	//	URL:          "www.Ajonob.com",
	//	CreationDate: "1401",
	//	UniversityID: 1,
	//	Courses: []Course{
	//		{ Name: "math", QuantityPlace: 1, StartDate: "1399", EndDate: "1402", CreatedDate: "1399", UniversityBranches: []UniversityBranch{{Model: gorm.Model{ID: 1}}} , Students: []Student{
	//			{ Name: "kasra", Age: 18, Mail: "kaskas@gmail.com", NationalCode: "015055", Address: "tehran", Courses: []Course{{Model: gorm.Model{ID: 1}}}, UniversityID: 1, University: University{} },
	//		}},
	//	},
	//}
	//teacher := Teacher{
	//	Name:         "alex",
	//	Mail:         "alex.mail",
	//	NationalCode: "075",
	//	Students:     []Student{
	//		{
	//			Model:gorm.Model{ID: 1},
	//		},
	//	},
	//}
	//err = db.Save(&uni).Error
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = db.Save(&uni_branch).Error
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	err = db.Save(&stu).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done ")
	return db
}

