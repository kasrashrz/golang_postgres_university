package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	viper.AutomaticEnv()
	viper_user := viper.Get("admin")
	viper_password := viper.Get("root")
	viper_db := viper.Get("go_university")
	viper_host := viper.Get("localhost")
	viper_port := viper.Get("36235")
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Students{})
	m := Students{Id: "ksa52ADj8ds2!", address: "teh"}
	db.Create(&m)
	return db
}
