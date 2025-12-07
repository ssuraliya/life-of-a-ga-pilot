package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Fleet{})
	db.AutoMigrate(&Flight{})

	return db
}
