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
		panic(err)
	}
	db = DB

	err = AutoMigrate(Product{}, CategoryLazada{})

	if err != nil {
		panic(err)
	}

}

func AutoMigrate(models ...interface{}) error {
	for idx := range models {
		if err := db.AutoMigrate(models[idx]); err != nil {
			return err
		}
	}
	return nil
}
