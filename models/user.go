package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	MiddleName string `json:"middle_name"`
	GroupID    uint   `json:"group_id" gorm:"default:1"` // будущая группа //дефолт - неопределившиеся
	Email      string `json:"email" gorm:"unique"`
	Phone      string `json:"phone" gorm:"unique"`
	IsDebtor   bool   `json:"is_debtor"`
	Password   string `json:"password"`
	Role       string `json:"role" gorm:"default:user"`
	//Внешние ключики
	Group Group `gorm:"foreignKey:GroupID;references:ID"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
