package db

import (
	"log"
	"os"

	"github.com/amine-bambrik-p8/go-lang-web-service/models"
	"github.com/jinzhu/gorm"

	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB represents a struct that contains information about the database
type DB struct {
	Database string
	Dialect  string
}

// Instance of Database information
var (
	Database DB
)

func init() {
	Database = DB{
		Dialect:  os.Getenv("DB_DIALECT"),
		Database: os.Getenv("DB_DATABASE"),
	}
}

// InitialMigration Migrates all the model tables
func InitialMigration() {
	database, err := gorm.Open(Database.Dialect, Database.Database)
	if err != nil {
		log.Print(err.Error())
		panic("Faild to connect to database")
	}
	defer database.Close()
	database.AutoMigrate(&models.User{})
}
