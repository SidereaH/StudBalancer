package models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
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

func (user *User) CreateUserWithoutDistrib(row []string) error {
	user.SecondName = row[0]
	user.FirstName = row[1]
	user.MiddleName = row[2]
	user.GroupID = 0
	user.Phone = row[3]
	user.Email = row[4]
	log.Println(fmt.Sprintf("Checking values: Specs - %s, %s; isDebtop - %s", row[6], row[7], row[8]))
	isDebtor := row[5]
	if isDebtor == "" || isDebtor == "false" || isDebtor == "нет" || isDebtor == "-" {
		user.IsDebtor = false
	} else if isDebtor == "true" || isDebtor == "+" || isDebtor == "да" {
		user.IsDebtor = true
	}
	if err := user.HashPassword(user.Phone); err != nil {
		return errors.New("Eror while creating pasword")
	}
	return nil
}
