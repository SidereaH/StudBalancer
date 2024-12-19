package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	GroupName      string `json:"group_name" gorm:"unique;not null"`
	SpecialityName string `json:"speciality_name" gorm:"unique;not null"`
	MaxSize        int    `json:"max_size" gorm:"default:30"`
}
