package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

  var (
	DB *gorm.DB
)
// This function initializes database connection
func InitDatabase() {
	const dsn = "host=localhost user=postgres password=secretpassword dbname=goats port=5432 sslmode=prefer TimeZone=Asia/Kolkata"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	if DB == nil {
		log.Println("Nil connection")
	}
	log.Println("Connection Opened to Database")
	fmt.Println("Database Migrated")
}
