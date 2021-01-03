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

	stu:=Student{
		Name:         "kasra",
		Age:          18,
		Mail:         "kasra@gmail",
		NationalCode: "015",
		Address:      "tehran",
		Courses:      []Course{
			{
				Model:gorm.Model{ID: 1},
			},
		},
		UniversityID: 1,
	}
	err = db.Save(&stu).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done ")
	return db
}

