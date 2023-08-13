package model

import (
	"fmt"
	"github.com/xuanvan229/crawl-golang/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable", config.Setting.DBUsername, config.Setting.DBPassword, config.Setting.DBHost, config.Setting.DBPort, config.Setting.DBName)
	DB, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		fmt.Println("hell oerror connect")
		panic(err)
	}
	db = DB
	fmt.Println("Connect to database successfully", db)

	err = db.AutoMigrate(Product{})

	if err != nil {
		fmt.Println("hello err migrate")
		panic(err)
	}
	fmt.Println("Auto migrate successfully")

}

func AutoMigrate(models ...interface{}) error {
	for idx := range models {
		if err := db.AutoMigrate(models[idx]); err != nil {
			return err
		}
	}
	return nil
}
