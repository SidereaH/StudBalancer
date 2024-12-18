package controllers

import (
	"net/http"
	"stud-distributor/database"
	"stud-distributor/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	//если хотим делать из телефона сразу пароль -
	//if err := user.HashPassword(user.Phone); err != nil {
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		//"userId": user.ID,
		"email":       user.Email,
		"phone":       user.Phone,
		"second_name": user.SecondName,
		"first_name":  user.FirstName,
		"middle_name": user.MiddleName,
		"group":       user.Group,
		"first_spec":  user.FirstSpec.Name,
		"second_spec": user.FirstSpec.Name})
}
