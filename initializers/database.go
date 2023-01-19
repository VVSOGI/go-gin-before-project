package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Database = DB

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	return Database
}
