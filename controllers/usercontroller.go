package controllers

import (
	"net/http"
	"stud-distributor/database"
	"stud-distributor/distributing"
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
	group, err := database.GetGroupByID(user.GroupID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		"group":       group})
}

type DistributeRequest struct {
	UserEmail      string `json:"user_mail"`
	FirstPriority  string `json:"first_priority"`
	SecondPriority string `json:"second_priority"`
}

func DistributeUser(context *gin.Context) {
	var request DistributeRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	//check if bro exists
	record := database.Instance.Where("email = ?", request.UserEmail).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	specs := []string{
		request.FirstPriority,
		request.SecondPriority,
	}
	if err := distributing.DistribureUserBySpecs(&user, specs); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	group, err := database.GetGroupByID(user.GroupID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"email": user.Email,
		"phone":       user.Phone,
		"second_name": user.SecondName,
		"first_name":  user.FirstName,
		"middle_name": user.MiddleName,
		"group":       group})
}
func GetUsers(c *gin.Context) {
	users, err := database.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := database.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

type EmailRequest struct {
	Email string `json:"user_email"`
}

func GetUserByEmail(c *gin.Context) {
	var request EmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	user, err := database.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"email": user.Email, "first_name": user.FirstName, "middle_name": user.MiddleName, "second_name": user.SecondName, "speciality_name": user.Group.SpecialityName, "group_name": user.Group.GroupName})

}
