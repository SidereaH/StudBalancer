package database

import (
	"errors"
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
		{
			GroupName:      "ИСП9-Kh31",
			SpecialityName: ".NET",
			MaxSize:        30,
		},
		{
			GroupName:      "ИСП9-Kh34",
			SpecialityName: "Data Engineer",
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
			log.Println("Error while creating group", result.Error)
		}
	}
	/*
		GroupName      string `json:"group_name" gorm:"unique"`
		SpecialityName string `json:"speciality_name" gorm:"unique"`
		MaxSize
	*/

}
func GetGroupByID(id uint) (models.Group, error) {
	var group models.Group
	err := Instance.Where("id = ?", id).Find(&group).Error
	if err != nil {
		return models.Group{}, err
	}
	return group, nil
}

// возврщает первую свободную группу по специальности
func GetGroupIdBySpecialityName(specialityName string) (uint, error) {
	var groups []models.Group
	err := Instance.Where("speciality_name = ?", specialityName).Find(&groups).Error
	if err != nil {
		return 0, err
	}
	for _, group := range groups {
		if isGroupIsFullByGroupID(group.ID) == false {
			return group.ID, nil
		}
	}
	return 0, fmt.Errorf("All groups is full now")
}
func isGroupIsFullByGroupID(groupID uint) bool {
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
func GetGroups() ([]models.Group, error) {
	var groups []models.Group
	if result := Instance.Find(&groups); result.Error != nil {
		errors.New("Failed to fetch groups")
		return nil, result.Error
	}
	return groups, nil
}
func DeleteGroupByID(id int) (bool, error) {
	result := Instance.Delete(&models.Group{}, id)
	if result.Error != nil {
		return false, fmt.Errorf("Failed to delete group with ID %d: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return false, fmt.Errorf("No group found with ID %d", id)
	}
	return true, nil
}
