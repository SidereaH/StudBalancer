package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stud-distributor/database"
	"stud-distributor/models"
)

type BulkCreateRequest struct {
	Groups []models.Group `json:"groups"`
}

func CreateGroups(c *gin.Context) {
	var req BulkCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(req.Groups) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No groups provided"})
		return
	}

	for _, group := range req.Groups {
		if err := database.Instance.Create(&group).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Groups created successfully"})
}
func GetGroups(c *gin.Context) {
	groups, err := database.GetGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"groups": groups})
}
func DeleteGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Преобразуем строку в int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = database.DeleteGroupByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully deleted group with id = %d", id)})
}
func GetGroupById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		c.Abort()
		return
	}
	group, err := database.GetGroupByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"group_name": group.GroupName, "group_size": group.MaxSize, "group_spec": group.SpecialityName})
}
func GetSpecialityNames(c *gin.Context) {
	specialities, err := database.GetUniqSpecialities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"specialities": specialities})
}
