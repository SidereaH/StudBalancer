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
	Instance.AutoMigrate(&models.User{}, &models.Spec{}, &models.Group{})
	InitSpecs(Instance)
	log.Println("Database Migration Completed!")
}

func InitSpecs(db *gorm.DB) {
	specs := []models.Spec{
		{
			Name: "Java",
		},
		{
			Name: "Backend",
		},
		{
			Name: "Frontend",
		},
		{
			Name: "Data Engineer",
		},
		{
			Name: ".NET",
		},
	}
	for _, spec := range specs {
		// Используем FirstOrCreate для предотвращения дублирования записей
		result := db.FirstOrCreate(&spec, models.Spec{Name: spec.Name})
		if result.Error != nil {
			log.Println("Ошибка при создании спецификации:", result.Error)
		}
	}

}
