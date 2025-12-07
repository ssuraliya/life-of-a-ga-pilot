package server

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	models "life-of-a-ga-pilot/internal/app/db"
)

func initDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Player{})
	db.AutoMigrate(&models.Fleet{})
	db.AutoMigrate(&models.Flight{})
}
