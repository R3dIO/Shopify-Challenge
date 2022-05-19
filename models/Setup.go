package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var dbname string = "local"

func ConnectDatabase() {
database, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{Logger: logger.Default.LogMode(logger.Error),})

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&Item{})

  DB = database
}