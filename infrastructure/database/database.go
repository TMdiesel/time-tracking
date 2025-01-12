package database

import (
	"log"
	"time-tracker/infrastructure/database/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("time-tracker.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&model.Project{}, &model.Task{})
	return db
}
