package models

import (
	"gorm.io/gorm"
)

type Spec struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);unique"` // Название специальности
}
