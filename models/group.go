package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	GroupName      string `json:"group_name" gorm:"unique"`
	SpecialityName string `json:"speciality_name" gorm:"unique"`
	MaxSize        int    `json:"max_size" `
}
