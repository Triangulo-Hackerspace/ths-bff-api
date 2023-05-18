package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Database struct
type Database struct {
	DB *gorm.DB
}

//NewDatabase : intializes and returns mysql db
func NewDatabase() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	DBPORT := os.Getenv("DB_PORT")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")

	URL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", HOST, USER, PASS, DBNAME, DBPORT, DB_SSLMODE)
	db, err := gorm.Open(postgres.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")

	}

	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}
