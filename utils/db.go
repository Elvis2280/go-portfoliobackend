package utils

import (
	Models "Portfoliobackend/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var database *gorm.DB

func ConnectDb() {
	connectionString := "host=127.0.0.1 user=admin password=admin dbname=portfolio port=5455 sslmode=disable"
	var error error
	db, error := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("applying migrations")

	err := db.AutoMigrate(&Models.Tag{}, &Models.Project{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(10)           // Set max open connections
	sqlDB.SetMaxIdleConns(10)           // Set max idle connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Set max connection lifetime to 1 hour
	fmt.Println("applied migrations")
	fmt.Println("Connected to Portfolio DB")
	database = db
}

func GetDb() (db *gorm.DB, err error) {
	sqlDb, err := database.DB()
	if err != nil {
		return database, err
	}

	if err := sqlDb.Ping(); err != nil {
		return database, err
	}
	return database, nil
}
