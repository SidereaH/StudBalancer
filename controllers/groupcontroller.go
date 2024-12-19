package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
