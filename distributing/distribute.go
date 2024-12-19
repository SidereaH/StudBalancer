package distributing

import (
	"errors"
	"stud-distributor/database"
	"stud-distributor/models"
)

func CreateUserWithoutDistrib(user *models.User, row []string) error {
	user.SecondName = row[0]
	user.FirstName = row[1]
	user.MiddleName = row[2]
	user.GroupID = 0
	user.Phone = row[3]
	user.Email = row[4]
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

func DistribureUserBySpecs(user *models.User, specs []string) error {
	var err error
	for _, spec := range specs {
		id, err := database.GetGroupIdBySpecialityName(spec)
		if err != nil {
			continue
		}
		user.GroupID = id
		return nil
	}
	return err
}
