package database

import (
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

		return nil, result.Error
	}
	for _, user := range users {
		group, err := GetGroupByID(user.GroupID)
		if err != nil {
			return nil, err
		}
		user.Group = group
		log.Println(user)
	}

	return users, nil
}
func GetUserByID(id string) (models.User, error) {
	var user models.User
	if result := Instance.First(&user, "id = ?", id); result.Error != nil {
		return models.User{}, result.Error
	}
	group, err := GetGroupByID(user.GroupID)
	if err != nil {
		return models.User{}, err
	}
	user.Group = group
	return user, nil
}
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if result := Instance.First(&user, "email = ?", email); result.Error != nil {
		return models.User{}, result.Error

	}
	group, err := GetGroupByID(user.GroupID)
	if err != nil {
		return models.User{}, err
	}
	user.Group = group
	return user, nil
}
