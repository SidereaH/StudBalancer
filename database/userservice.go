package database

import (
	"errors"
	"log"
	"stud-distributor/models"
)

func ExistsByEmail(record []string) bool {
	var user models.User
	recording := Instance.Where("email = ?", record[5]).First(&user)
	if recording.Error != nil {
		//не найден - найс, регаем
		log.Println(recording.Error)
		return false
	}
	return true
}
func ExistsByPhone(record []string) bool {
	var user models.User
	recording := Instance.Where("phone = ?", record[4]).First(&user)
	if recording.Error != nil {
		//не найден - найс, регаем
		log.Println(recording.Error)
		return false
	}
	return true
}
func GetUsers() ([]models.User, error) {
	var users []models.User
	if result := Instance.Find(&users); result.Error != nil {
		errors.New("Failed to fetch users")
		return nil, result.Error
	}
	return users, nil
}
