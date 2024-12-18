package models

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"stud-distributor/database"
)

type User struct {
	gorm.Model
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	MiddleName   string `json:"middle_name"`
	Group        string `json:"group"`
	FirstSpecID  int8   `json:"firstspec_id"`
	SecondSpecID int    `json:"secondspec_id"`
	Email        string `json:"email" gorm:"unique"`
	Phone        string `json:"phone" gorm:"unique"`
	IsDebtor     bool   `json:"is_debtor"`
	Password     string `json:"password"`

	//Внешние ключики
	FirstSpec  Spec `gorm:"foreignKey:FirstSpecID;references:ID"`
	SecondSpec Spec `gorm:"foreignKey:SecondSpecID;references:ID"`
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

func (user *User) FillUserParams(row []string) error {
	user.SecondName = row[0]
	user.FirstName = row[1]
	user.MiddleName = row[2]
	user.Group = row[3]
	user.Phone = row[4]
	user.Email = row[5]

	log.Println(fmt.Sprintf("Checking values: Specs - %s, %s; isDebtop - %s", row[6], row[7], row[8]))

	if row[6] == "" {

	} else {
		spec1, err := getSpecIdByName(row[6])
		if err != nil {
			log.Println(err)
			return err
		}
		user.FirstSpecID = int8(spec1)
	}
	if row[7] == "" {
		user.SecondSpec.Name = "Не определился"
	} else {
		user.SecondSpec.Name = row[7]
	}
	if row[8] == "" || row[8] == "false" || row[8] == "нет" || row[8] == "-" {
		user.IsDebtor = false
	} else if row[8] == "true" || row[8] == "+" || row[8] == "да" {
		user.IsDebtor = true
	}
	if err := user.HashPassword(user.Phone); err != nil {
		return errors.New("Eror while creating pasword")
	}
	return nil
}
func getSpecIdByName(name string) (uint, error) {
	var spec Spec
	record := database.Instance.Where("name = ?", name).First(&spec)
	if record.Error != nil {
		log.Println(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return 0, errors.New("error" + record.Error.Error())
	}
	return spec.ID, nil
}
