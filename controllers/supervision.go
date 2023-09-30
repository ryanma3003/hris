package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func SupervisionIndex(c *gin.Context) {
	var sup []models.Supervision
	err := db.DB.Preload("Division").Preload("Department").Find(&sup).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": sup,
	})
}

func SupervisionCreate(c *gin.Context) {
	var body models.Supervision

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	sup := models.Supervision{Name: body.Name, DivisionID: body.DivisionID, DepartmentID: body.DepartmentID}
	result := db.DB.Create(&sup)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"message": "create success",
	})
}

func SupervisionShow(c *gin.Context) {
	id := c.Param("id")

	var sup models.Supervision
	err := db.DB.First(&sup, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": sup,
	})
}

func SupervisionUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Supervision

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var sup models.Supervision
	err := db.DB.First(&sup, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&sup).Updates(models.Supervision{
		Name:         body.Name,
		DivisionID:   body.DivisionID,
		DepartmentID: body.DepartmentID,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func SupervisionDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var sup models.Supervision
	err := db.DB.Delete(&sup, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}
