
package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "time-tracker/infrastructure/database/model"
)

func NewDatabase() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("time-tracker.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    db.AutoMigrate(&model.Project{})
    return db
}
