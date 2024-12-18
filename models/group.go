package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName string `json:"group_name" `
	User      User   `json:"email" ` // Электронная почта
}
