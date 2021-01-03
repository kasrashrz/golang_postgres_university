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
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,
	//})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&University{},&UniversityBranch{},&Student{},&Course{},&Teacher{})
	//uni := UniversityBranch{
	//	Name:         "azad shomal",
	//	Address:      "tehran north",
	//	URL:          "www.northAzad.ir",
	//	CreationDate: "1400",
	//	UniversityID: 1,
	//	Courses:      nil,
	//}
	//crs := Course{
	//	Name:               "math",
	//	QuantityPlace:      2,
	//	StartDate:          "1400",
	//	EndDate:            "1400",
	//	CreatedDate:        "1400",
	//	TeacherID:          1,
	//	UniversityBranches: []UniversityBranch{{
	//		Model:gorm.Model{ID: 1},
	//	}},
	//	Students:           nil,
	//}
	//tea := Teacher{
	//	Name:         "gholam",
	//	Mail:         "gholi@gmail.com",
	//	NationalCode: "015055",
	//}
	stu := Student{
		Name:         "jafar",
		Age:          15,
		Mail:         "jefri@gmail.com",
		NationalCode: "00015",
		Address:      "boshehr",
		Courses:      []Course{
			{
				Model:gorm.Model{ID: 5},
			},
		},
		UniversityID: 1,
	}

	err = db.Save(&stu).Error
	if err != nil{
		panic(err)
	}

	stu1 := Student{
		Name:         "kasra",
		Age:          18,
		Mail:         "kas@mail.com",
		NationalCode: "0150553692",
		Address:      "teh",
		Courses: []Course{
			{
				Model:gorm.Model{ID: 5},
			},
		},
		UniversityID: 1,
	}
	err = db.Save(&stu1).Error
	if err != nil{
		panic(err)
	}

	fmt.Println("Done ")
	return db
}

