package database

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"stud-distributor/models"
)

func InitGroups(db *gorm.DB) {
	groups := []models.Group{
		{
			GroupName:      "Не определился",
			SpecialityName: "Не определился",
			MaxSize:        0,
		},
		{
			GroupName:      "ИСП11-Kh21",
			SpecialityName: "Backend",
			MaxSize:        30,
		},
		{
			GroupName:      "ИСП9-Kh33",
			SpecialityName: "Java",
			MaxSize:        30,
		},
		{
			GroupName:      "ИСП9-Kh32",
			SpecialityName: "Frontend",
			MaxSize:        30,
		},
	}
	for _, group := range groups {
		// Используем FirstOrCreate для предотвращения дублирования записей
		result := db.FirstOrCreate(&group, models.Group{
			GroupName:      group.GroupName,
			SpecialityName: group.SpecialityName,
			MaxSize:        group.MaxSize})
		if result.Error != nil {
			log.Println("Ошибка при создании спецификации:", result.Error)
		}
	}
	/*
		GroupName      string `json:"group_name" gorm:"unique"`
		SpecialityName string `json:"speciality_name" gorm:"unique"`
		MaxSize
	*/

}
func GetGroupByID(id uint) ([]models.Group, error) {
	var group []models.Group
	err := Instance.Where("id = ?", id).Find(&group).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}
func GetGroupsBySpecialityName(specialityName string) ([]models.Group, error) {
	var groups []models.Group
	err := Instance.Where("speciality_name = ?", specialityName).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}
func IsGroupIsFullByGroupID(groupID uint) bool {
	var group models.Group
	Instance.First(&group, "ID = ?", groupID)
	groupSize := group.MaxSize
	currentSize := GetCountOfUsersByGroupID(group.ID)
	if currentSize < groupSize {
		return false
	}
	return true
}

func GetCountOfUsersByGroupID(groupID uint) int {
	var count int64
	err := Instance.Model(&models.User{}).Where("group_id = ?", groupID).Count(&count).Error
	if err != nil {
		// Обработать ошибку, например, вывести в лог
		fmt.Println("Error counting users:", err)
		return 0
	}
	return int(count)
}
