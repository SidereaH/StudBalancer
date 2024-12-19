package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"stud-distributor/models"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{}, &models.Group{})
	InitGroups(Instance)
	log.Println("Database Migration Completed!")
}
